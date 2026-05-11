package binlog

import (
	"context"
	"database/sql"
	"fmt"
	"shop-goframe-micro-service-refacotor/app/search/utility/elasticsearch"
	"time"

	"github.com/go-mysql-org/go-mysql/mysql"

	"github.com/go-mysql-org/go-mysql/replication"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/olivere/elastic/v7"
)

var goodsTimeLocation = loadGoodsTimeLocation()

// SyncAllGoodsToES 全量同步 goods.goods_info 到 ES。
func SyncAllGoodsToES(ctx context.Context) error {
	db, err := sql.Open("mysql", mysqlDSN(ctx))
	if err != nil {
		return fmt.Errorf("打开 MySQL 连接失败: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("连接 MySQL 失败: %w", err)
	}

	rows, err := db.QueryContext(ctx, `
		SELECT id, name, images, price, level1_category_id, level2_category_id,
		       level3_category_id, brand, stock, sale, tags, detail_info,
		       created_at, updated_at, deleted_at
		FROM goods.goods_info
	`)
	if err != nil {
		return fmt.Errorf("查询 goods.goods_info 失败: %w", err)
	}
	defer rows.Close()

	ids := make(map[string]struct{})
	var count int
	for rows.Next() {
		data, err := scanGoodsRow(rows)
		if err != nil {
			return err
		}
		if ok := upsertToES(ctx, data); ok {
			ids[gconv.String(data["id"])] = struct{}{}
			count++
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("读取 goods.goods_info 失败: %w", err)
	}
	if err := deleteStaleGoodsFromES(ctx, ids); err != nil {
		return err
	}

	g.Log().Infof(ctx, "全量同步商品到ES完成: count=%d", count)
	return nil
}

// StartBinlogSyncer 启动 MySQL Binlog 同步
func StartBinlogSyncer(ctx context.Context) {
	// 从配置获取 MySQL 连接信息
	mysqlHost := g.Cfg().MustGet(ctx, "binlog.goods.mysql.host").String()
	mysqlPort := g.Cfg().MustGet(ctx, "binlog.goods.mysql.port").Int()
	mysqlUser := g.Cfg().MustGet(ctx, "binlog.goods.mysql.username").String()
	mysqlPass := g.Cfg().MustGet(ctx, "binlog.goods.mysql.password").String()

	// 创建 Binlog 同步器
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100, // 需要唯一
		Flavor:   "mysql",
		Host:     mysqlHost,
		Port:     uint16(mysqlPort),
		User:     mysqlUser,
		Password: mysqlPass,
	}

	syncer := replication.NewBinlogSyncer(cfg)
	defer syncer.Close()

	position, err := currentBinlogPosition(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取当前 Binlog 位点失败: %v", err)
		return
	}

	// 开始同步
	streamer, err := syncer.StartSync(position)
	if err != nil {
		g.Log().Errorf(ctx, "启动 Binlog 同步失败: %v", err)
		return
	}

	g.Log().Info(ctx, "开始监听 MySQL Binlog...")

	for {
		// 获取 Binlog 事件
		ev, err := streamer.GetEvent(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "获取 Binlog 事件失败: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// 处理事件
		processBinlogEvent(ctx, ev)
	}
}

// processBinlogEvent 处理 Binlog 事件
func processBinlogEvent(ctx context.Context, ev *replication.BinlogEvent) {
	switch e := ev.Event.(type) {
	case *replication.RowsEvent:
		// 只处理指定数据库的 goods_info 表
		if string(e.Table.Schema) != "goods" || string(e.Table.Table) != "goods_info" {
			return
		}

		g.Log().Debugf(ctx, "收到Binlog事件: 数据库=%s, 表=%s", e.Table.Schema, e.Table.Table)

		// 根据事件类型处理
		switch ev.Header.EventType {
		case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2:
			handleInsert(ctx, e.Rows)
		case replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2:
			handleUpdate(ctx, e.Rows)
		case replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
			handleDelete(ctx, e.Rows)
		default:
		}
	}
}

// handleInsert 处理插入事件
func handleInsert(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		// 将行数据转换为 map
		columnMap := parseRowData(row)
		upsertToES(ctx, columnMap)
	}
}

// handleUpdate 处理更新事件
func handleUpdate(ctx context.Context, rows [][]interface{}) {
	// 更新事件的行数据格式为 [旧行数据, 新行数据]
	for i := 0; i < len(rows); i += 2 {
		if i+1 < len(rows) {
			columnMap := parseRowData(rows[i+1]) // 取新数据
			upsertToES(ctx, columnMap)
		}
	}
}

// handleDelete 处理删除事件
func handleDelete(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		columnMap := parseRowData(row)
		deleteFromES(ctx, columnMap)
	}
}

// parseRowData 解析行数据为 map
func parseRowData(row []interface{}) map[string]interface{} {
	// 这里需要根据你的表结构定义字段名
	fields := []string{
		"id", "name", "images", "price", "level1_category_id",
		"level2_category_id", "level3_category_id", "brand",
		"stock", "sale", "tags", "detail_info",
		"created_at", "updated_at", "deleted_at",
	}

	result := make(map[string]interface{})
	for i, value := range row {
		if i < len(fields) {
			result[fields[i]] = value
		}
	}
	return result
}

// upsertToES 插入或更新文档到 ES
func upsertToES(ctx context.Context, data map[string]interface{}) bool {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return false
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	doc := buildGoodsDocument(data)
	id := gconv.String(doc["id"])

	_, err := client.Index().
		Index(esIndexGoods).
		Id(id).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		g.Log().Errorf(ctx, "同步商品到ES失败: %v", err)
		return false
	} else {
		g.Log().Debugf(ctx, "成功同步商品到ES: ID=%s", id)
		return true
	}
}

// deleteFromES 从 ES 删除文档
func deleteFromES(ctx context.Context, data map[string]interface{}) {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return
	}

	id := gconv.String(data["id"])
	if id == "" {
		g.Log().Error(ctx, "删除操作未找到ID")
		return
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()

	_, err := client.Delete().
		Index(esIndexGoods).
		Id(id).
		Do(ctx)

	if err != nil {
		g.Log().Errorf(ctx, "从ES删除商品失败: ID=%s, error=%v", id, err)
	} else {
		g.Log().Debugf(ctx, "成功从ES删除商品: ID=%s", id)
	}
}

func currentBinlogPosition(ctx context.Context) (mysql.Position, error) {
	db, err := sql.Open("mysql", mysqlDSN(ctx))
	if err != nil {
		return mysql.Position{}, err
	}
	defer db.Close()

	var (
		file           string
		position       uint32
		binlogDoDB     sql.NullString
		binlogIgnoreDB sql.NullString
		executedGTID   sql.NullString
	)
	if err := db.QueryRowContext(ctx, "SHOW MASTER STATUS").Scan(
		&file,
		&position,
		&binlogDoDB,
		&binlogIgnoreDB,
		&executedGTID,
	); err != nil {
		return mysql.Position{}, err
	}

	return mysql.Position{Name: file, Pos: position}, nil
}

func deleteStaleGoodsFromES(ctx context.Context, validIDs map[string]struct{}) error {
	client := elasticsearch.GetClient()
	if client == nil {
		return fmt.Errorf("ES客户端未初始化")
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	result, err := client.Search().
		Index(esIndexGoods).
		Query(elastic.NewMatchAllQuery()).
		FetchSource(false).
		Size(10000).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("查询ES旧商品失败: %w", err)
	}

	for _, hit := range result.Hits.Hits {
		if _, ok := validIDs[hit.Id]; ok {
			continue
		}
		if _, err := client.Delete().Index(esIndexGoods).Id(hit.Id).Do(ctx); err != nil {
			return fmt.Errorf("删除ES旧商品失败: ID=%s, error=%w", hit.Id, err)
		}
		g.Log().Debugf(ctx, "删除ES旧商品: ID=%s", hit.Id)
	}
	return nil
}

func mysqlDSN(ctx context.Context) string {
	mysqlHost := g.Cfg().MustGet(ctx, "binlog.goods.mysql.host").String()
	mysqlPort := g.Cfg().MustGet(ctx, "binlog.goods.mysql.port").Int()
	mysqlUser := g.Cfg().MustGet(ctx, "binlog.goods.mysql.username").String()
	mysqlPass := g.Cfg().MustGet(ctx, "binlog.goods.mysql.password").String()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/goods?parseTime=true&loc=Asia%%2FShanghai", mysqlUser, mysqlPass, mysqlHost, mysqlPort)
}

func scanGoodsRow(rows *sql.Rows) (map[string]interface{}, error) {
	var (
		id               uint32
		name             string
		images           sql.NullString
		price            uint64
		level1CategoryId uint32
		level2CategoryId uint32
		level3CategoryId uint32
		brand            string
		stock            uint32
		sale             uint32
		tags             string
		detailInfo       sql.NullString
		createdAt        sql.NullTime
		updatedAt        sql.NullTime
		deletedAt        sql.NullTime
	)

	if err := rows.Scan(
		&id, &name, &images, &price, &level1CategoryId, &level2CategoryId,
		&level3CategoryId, &brand, &stock, &sale, &tags, &detailInfo,
		&createdAt, &updatedAt, &deletedAt,
	); err != nil {
		return nil, fmt.Errorf("解析 goods.goods_info 行失败: %w", err)
	}

	return map[string]interface{}{
		"id":                 id,
		"name":               name,
		"images":             nullStringValue(images),
		"price":              price,
		"level1_category_id": level1CategoryId,
		"level2_category_id": level2CategoryId,
		"level3_category_id": level3CategoryId,
		"brand":              brand,
		"stock":              stock,
		"sale":               sale,
		"tags":               tags,
		"detail_info":        nullStringValue(detailInfo),
		"created_at":         createdAt,
		"updated_at":         updatedAt,
		"deleted_at":         deletedAt,
	}, nil
}

func buildGoodsDocument(data map[string]interface{}) map[string]interface{} {
	doc := map[string]interface{}{
		"id":                 gconv.Uint32(data["id"]),
		"name":               gconv.String(data["name"]),
		"images":             normalizeImages(data["images"]),
		"price":              gconv.Uint64(data["price"]),
		"level1_category_id": gconv.Uint32(data["level1_category_id"]),
		"level2_category_id": gconv.Uint32(data["level2_category_id"]),
		"level3_category_id": gconv.Uint32(data["level3_category_id"]),
		"brand":              gconv.String(data["brand"]),
		"stock":              gconv.Uint32(data["stock"]),
		"sale":               gconv.Uint32(data["sale"]),
		"tags":               gconv.String(data["tags"]),
		"detail_info":        gconv.String(data["detail_info"]),
	}

	if value, ok := normalizeESTime(data["created_at"]); ok {
		doc["created_at"] = value
	}
	if value, ok := normalizeESTime(data["updated_at"]); ok {
		doc["updated_at"] = value
	}
	if value, ok := normalizeESTime(data["deleted_at"]); ok {
		doc["deleted_at"] = value
	}

	return doc
}

func normalizeImages(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return ""
	case []byte:
		return string(v)
	case sql.NullString:
		return nullStringValue(v)
	default:
		return gconv.String(v)
	}
}

func normalizeESTime(value interface{}) (string, bool) {
	switch v := value.(type) {
	case nil:
		return "", false
	case time.Time:
		return formatESTime(v)
	case *time.Time:
		if v == nil {
			return "", false
		}
		return formatESTime(*v)
	case sql.NullTime:
		if !v.Valid {
			return "", false
		}
		return formatESTime(v.Time)
	case []byte:
		return parseESTimeString(string(v))
	case string:
		return parseESTimeString(v)
	default:
		return parseESTimeString(gconv.String(v))
	}
}

func formatESTime(t time.Time) (string, bool) {
	if t.IsZero() {
		return "", false
	}
	return t.In(goodsTimeLocation).Format(time.RFC3339), true
}

func parseESTimeString(value string) (string, bool) {
	if value == "" || value == "0000-00-00 00:00:00" || value == "<nil>" {
		return "", false
	}

	layouts := []string{
		time.RFC3339Nano,
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
	}
	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, goodsTimeLocation); err == nil {
			return formatESTime(t)
		}
	}
	return value, true
}

func nullStringValue(value sql.NullString) string {
	if !value.Valid {
		return ""
	}
	return value.String
}

func loadGoodsTimeLocation() *time.Location {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Local
	}
	return location
}

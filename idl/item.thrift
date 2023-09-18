struct Item {
    1: i64 ID,
    2: string Name,
}

struct GetItemRequest {
    1: i64 ID (api.query="id")
}

struct GetItemResponse {
    1: optional Item Item,
}

service ItemService {
    GetItemResponse GetItem(1: GetItemRequest req) (api.get="/item")
}
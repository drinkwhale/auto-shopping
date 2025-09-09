package graph

import "github.com/drinkwhale/auto-shopping/internal/graph/model"

// Resolver 구조체는 앱의 의존성을 관리합니다.
// 여기에 todos 슬라이스를 추가하여 임시 데이터베이스로 사용합니다.
type Resolver struct {
	todos []*model.Todo
}

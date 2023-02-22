// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/interfaces.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"
	usecase "softpro6/internal/usecase"
	valueobject "softpro6/internal/valueobject"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockGetLineUseCase is a mock of GetLineUseCase interface.
type MockGetLineUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockGetLineUseCaseMockRecorder
}

// MockGetLineUseCaseMockRecorder is the mock recorder for MockGetLineUseCase.
type MockGetLineUseCaseMockRecorder struct {
	mock *MockGetLineUseCase
}

// NewMockGetLineUseCase creates a new mock instance.
func NewMockGetLineUseCase(ctrl *gomock.Controller) *MockGetLineUseCase {
	mock := &MockGetLineUseCase{ctrl: ctrl}
	mock.recorder = &MockGetLineUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetLineUseCase) EXPECT() *MockGetLineUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetLineUseCase) Execute(ctx context.Context, sport string) (usecase.Line, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, sport)
	ret0, _ := ret[0].(usecase.Line)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetLineUseCaseMockRecorder) Execute(ctx, sport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetLineUseCase)(nil).Execute), ctx, sport)
}

// MockGetLineProvider is a mock of GetLineProvider interface.
type MockGetLineProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGetLineProviderMockRecorder
}

// MockGetLineProviderMockRecorder is the mock recorder for MockGetLineProvider.
type MockGetLineProviderMockRecorder struct {
	mock *MockGetLineProvider
}

// NewMockGetLineProvider creates a new mock instance.
func NewMockGetLineProvider(ctrl *gomock.Controller) *MockGetLineProvider {
	mock := &MockGetLineProvider{ctrl: ctrl}
	mock.recorder = &MockGetLineProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetLineProvider) EXPECT() *MockGetLineProviderMockRecorder {
	return m.recorder
}

// GetLine mocks base method.
func (m *MockGetLineProvider) GetLine(ctx context.Context, sportName string) (usecase.Line, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLine", ctx, sportName)
	ret0, _ := ret[0].(usecase.Line)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLine indicates an expected call of GetLine.
func (mr *MockGetLineProviderMockRecorder) GetLine(ctx, sportName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLine", reflect.TypeOf((*MockGetLineProvider)(nil).GetLine), ctx, sportName)
}

// MockLine is a mock of Line interface.
type MockLine struct {
	ctrl     *gomock.Controller
	recorder *MockLineMockRecorder
}

// MockLineMockRecorder is the mock recorder for MockLine.
type MockLineMockRecorder struct {
	mock *MockLine
}

// NewMockLine creates a new mock instance.
func NewMockLine(ctrl *gomock.Controller) *MockLine {
	mock := &MockLine{ctrl: ctrl}
	mock.recorder = &MockLineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLine) EXPECT() *MockLineMockRecorder {
	return m.recorder
}

// Rate mocks base method.
func (m *MockLine) Rate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rate")
	ret0, _ := ret[0].(string)
	return ret0
}

// Rate indicates an expected call of Rate.
func (mr *MockLineMockRecorder) Rate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rate", reflect.TypeOf((*MockLine)(nil).Rate))
}

// Sport mocks base method.
func (m *MockLine) Sport() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sport")
	ret0, _ := ret[0].(string)
	return ret0
}

// Sport indicates an expected call of Sport.
func (mr *MockLineMockRecorder) Sport() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sport", reflect.TypeOf((*MockLine)(nil).Sport))
}

// MockPollProviderUseCase is a mock of PollProviderUseCase interface.
type MockPollProviderUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockPollProviderUseCaseMockRecorder
}

// MockPollProviderUseCaseMockRecorder is the mock recorder for MockPollProviderUseCase.
type MockPollProviderUseCaseMockRecorder struct {
	mock *MockPollProviderUseCase
}

// NewMockPollProviderUseCase creates a new mock instance.
func NewMockPollProviderUseCase(ctrl *gomock.Controller) *MockPollProviderUseCase {
	mock := &MockPollProviderUseCase{ctrl: ctrl}
	mock.recorder = &MockPollProviderUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPollProviderUseCase) EXPECT() *MockPollProviderUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockPollProviderUseCase) Execute(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockPollProviderUseCaseMockRecorder) Execute(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockPollProviderUseCase)(nil).Execute), ctx)
}

// MockStoreSportUseCase is a mock of StoreSportUseCase interface.
type MockStoreSportUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockStoreSportUseCaseMockRecorder
}

// MockStoreSportUseCaseMockRecorder is the mock recorder for MockStoreSportUseCase.
type MockStoreSportUseCaseMockRecorder struct {
	mock *MockStoreSportUseCase
}

// NewMockStoreSportUseCase creates a new mock instance.
func NewMockStoreSportUseCase(ctrl *gomock.Controller) *MockStoreSportUseCase {
	mock := &MockStoreSportUseCase{ctrl: ctrl}
	mock.recorder = &MockStoreSportUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreSportUseCase) EXPECT() *MockStoreSportUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockStoreSportUseCase) Execute(ctx context.Context, sport usecase.Sport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", ctx, sport)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockStoreSportUseCaseMockRecorder) Execute(ctx, sport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockStoreSportUseCase)(nil).Execute), ctx, sport)
}

// MockSport is a mock of Sport interface.
type MockSport struct {
	ctrl     *gomock.Controller
	recorder *MockSportMockRecorder
}

// MockSportMockRecorder is the mock recorder for MockSport.
type MockSportMockRecorder struct {
	mock *MockSport
}

// NewMockSport creates a new mock instance.
func NewMockSport(ctrl *gomock.Controller) *MockSport {
	mock := &MockSport{ctrl: ctrl}
	mock.recorder = &MockSportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSport) EXPECT() *MockSportMockRecorder {
	return m.recorder
}

// CreatedAt mocks base method.
func (m *MockSport) CreatedAt() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// CreatedAt indicates an expected call of CreatedAt.
func (mr *MockSportMockRecorder) CreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockSport)(nil).CreatedAt))
}

// Name mocks base method.
func (m *MockSport) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockSportMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockSport)(nil).Name))
}

// Rate mocks base method.
func (m *MockSport) Rate() valueobject.Rate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rate")
	ret0, _ := ret[0].(valueobject.Rate)
	return ret0
}

// Rate indicates an expected call of Rate.
func (mr *MockSportMockRecorder) Rate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rate", reflect.TypeOf((*MockSport)(nil).Rate))
}

// MockSportRepository is a mock of SportRepository interface.
type MockSportRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSportRepositoryMockRecorder
}

// MockSportRepositoryMockRecorder is the mock recorder for MockSportRepository.
type MockSportRepositoryMockRecorder struct {
	mock *MockSportRepository
}

// NewMockSportRepository creates a new mock instance.
func NewMockSportRepository(ctrl *gomock.Controller) *MockSportRepository {
	mock := &MockSportRepository{ctrl: ctrl}
	mock.recorder = &MockSportRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSportRepository) EXPECT() *MockSportRepositoryMockRecorder {
	return m.recorder
}

// GetRecent mocks base method.
func (m *MockSportRepository) GetRecent(ctx context.Context) (usecase.Sport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecent", ctx)
	ret0, _ := ret[0].(usecase.Sport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecent indicates an expected call of GetRecent.
func (mr *MockSportRepositoryMockRecorder) GetRecent(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecent", reflect.TypeOf((*MockSportRepository)(nil).GetRecent), ctx)
}

// IsSynced mocks base method.
func (m *MockSportRepository) IsSynced(ctx context.Context, after time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSynced", ctx, after)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSynced indicates an expected call of IsSynced.
func (mr *MockSportRepositoryMockRecorder) IsSynced(ctx, after interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSynced", reflect.TypeOf((*MockSportRepository)(nil).IsSynced), ctx, after)
}

// Store mocks base method.
func (m *MockSportRepository) Store(ctx context.Context, sport usecase.Sport) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, sport)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockSportRepositoryMockRecorder) Store(ctx, sport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockSportRepository)(nil).Store), ctx, sport)
}

// MockLineToSportPolicy is a mock of LineToSportPolicy interface.
type MockLineToSportPolicy struct {
	ctrl     *gomock.Controller
	recorder *MockLineToSportPolicyMockRecorder
}

// MockLineToSportPolicyMockRecorder is the mock recorder for MockLineToSportPolicy.
type MockLineToSportPolicyMockRecorder struct {
	mock *MockLineToSportPolicy
}

// NewMockLineToSportPolicy creates a new mock instance.
func NewMockLineToSportPolicy(ctrl *gomock.Controller) *MockLineToSportPolicy {
	mock := &MockLineToSportPolicy{ctrl: ctrl}
	mock.recorder = &MockLineToSportPolicyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLineToSportPolicy) EXPECT() *MockLineToSportPolicyMockRecorder {
	return m.recorder
}

// Export mocks base method.
func (m *MockLineToSportPolicy) Export(arg0 usecase.Line) (usecase.Sport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Export", arg0)
	ret0, _ := ret[0].(usecase.Sport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Export indicates an expected call of Export.
func (mr *MockLineToSportPolicyMockRecorder) Export(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Export", reflect.TypeOf((*MockLineToSportPolicy)(nil).Export), arg0)
}
package bootstrap

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)
// Mock the actual database connector
const (
	ENV_FILE = "../../.env"
)
type MockMongoClientManager struct {
	mock.Mock
}

func (m *MockMongoClientManager) Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	args := m.Called(ctx, uri)
	return args.Get(0).(*mongo.Client), args.Error(1)
}

func (m *MockMongoClientManager) Ping(ctx context.Context, client *mongo.Client) error {
	args := m.Called(ctx, client)
	return args.Error(0)
}

func (m *MockMongoClientManager) Disconnect(client *mongo.Client) error{
	args := m.Called(client)
	return args.Error(0)
}

// Database successfully connects 
func TestNewMongoDatabase_Success(t *testing.T) {
	mmcm := new(MockMongoClientManager)
	mock_client := &mongo.Client{}
	mock_env, _ := NewEnv(ENV_FILE)
	mock_env.MongoUri = "mongodb://mock-uri"

	mmcm.On("Connect", mock.Anything, mock_env.MongoUri).Return(mock_client, nil)
	mmcm.On("Ping", mock.Anything, mock_client).Return(nil)


	client, err := NewMongoDatabase(mock_env, mmcm)
	assert.NoError(t, err, "Expect no error")
	assert.Equal(t, mock_client, client, "Expect client")
	mmcm.AssertExpectations(t)
}

// Connect fails due to bad uri
func TestNewMongoDatabase_ConnectFails(t *testing.T) {
	mmcm := new(MockMongoClientManager)
	env, _ := NewEnv(ENV_FILE)
	env.MongoUri = "mongodb://bad-uri"

	mmcm.On("Connect", mock.Anything, env.MongoUri).Return((*mongo.Client)(nil), errors.New("connection failed"))

	client, err := NewMongoDatabase(env, mmcm)

	assert.Nil(t, client)
	assert.EqualError(t, err, "connection failed")
	mmcm.AssertExpectations(t)

}

// Pinging fails
func TestNewMongoDatabase_PingFails(t *testing.T) {
	mmcm := new(MockMongoClientManager)
	env, _ := NewEnv(ENV_FILE)
	env.MongoUri = "mongodb://mock-uri"
	mockClient := &mongo.Client{}

	mmcm.
		On("Connect", mock.Anything, env.MongoUri).
		Return(mockClient, nil)
	mmcm.
		On("Ping", mock.Anything, mockClient).
		Return(errors.New("ping failed"))

	client, err := NewMongoDatabase(env, mmcm)

	assert.Nil(t, client)
	assert.EqualError(t, err, "ping failed")
	mmcm.AssertExpectations(t)
}


// Disconnect succussful

func TestDisconnect_Success(t *testing.T) {
	mmcm := new(MockMongoClientManager)
	mockClient := &mongo.Client{}

	mmcm.On("Disconnect", mockClient).Return(nil)

	err := mmcm.Disconnect(mockClient)
	assert.NoError(t, err, "Expect no error on successful disconnect")
	mmcm.AssertExpectations(t)
}

// Disconnect fails
func TestDisconnect_Fails(t *testing.T) {
	mmcm := new(MockMongoClientManager)
	mockClient := &mongo.Client{}

	mmcm.On("Disconnect", mockClient).Return(errors.New("disconnect failed"))

	err := mmcm.Disconnect(mockClient)
	assert.EqualError(t, err, "disconnect failed")
	mmcm.AssertExpectations(t)
}
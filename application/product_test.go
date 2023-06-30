package application_test

import( 
	"testing"
	"github.com/streechr/testify/require"
)

func TestProduct_Enable(t *testing.T){
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err:= product.Enable()
	require.Nil(t, err)
	
	product.Price = 0
	err:= product.Enable()
	require.Equal(t, "Invalid price; Must be greather than zero", err.Error)
}


func TestProduct_Disable(t *testing.T){
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 10

	err:= product.Disable()
	require.Equal(t, "Invalid price; Must be zero", err.Error)

	
	product.Price = 0
	err:= product.Disable()
	require.Nil(t, err)	
}

func TestProduct_IsValid(t *testing.T){
	product:= application.Product{}
	product.Name = "Hello"
	product.ID = uuid.NewV4().String()
	product.Status = application.ENABLED
	product.Price = 10

	_, err =: product.IsValid()
	required.Nil(t, err)

	product.Status = "INVALID"
	_, err =: product.IsValid()
	require.Equal(t, "Invalid status", err.Error)

}
/**
 * Master Controller
 */
angular.module('Rest')
    .controller('OrderController', ['$scope', 'localStorageService', 'STATUS', OrderController]);

function OrderController( $scope, localStorageService, STATUS ) {		  
    $scope.clearAll = function(){
      _.each( $scope.orders, function( obj ){
        obj.quantity = 0;
      });
      
      $scope.orders.splice( 0 );
    }    

    $scope.plus = function( meal ){ 

        $scope.ordersClass = 'active';

        if( !_.findWhere( $scope.orders ,{ id: meal.id } ) )
            $scope.orders.push( meal );

        meal.quantity += 1;
    };

    $scope.minus = function( meal ){        

        if( meal.quantity != 0 )
        {   
            meal.quantity -= 1;            
        } 
        if( meal.quantity == 0 ){
            $scope.orders.splice( _.indexOf( $scope.orders, meal ), 1 );
        }        
        if( !$scope.orders.length )
        {
            $scope.ordersClass = '';
            console.log( $scope.orders.length );
        }
    };

    $scope.order = function(){

        var orders = localStorageService.get('orders') || [];

        var order = {
            id: 2,
            status: STATUS.ISSUED,         
            items: $scope.orders,
            tableId: localStorageService.get('tableId'),
            cost: $scope.getTotalCost(),            
        };

        orders.push( order );

        localStorageService.set('orders', orders);

        $scope.clearAll();

        $scope.sentOrders = $scope.getOrders();
        
    };    

    $scope.totalCost = 0;

    $scope.status = [ '', 'issued', 'accepted', 'inprogress', 'ready', 'delivered', 'paid', 'cancelled' ];
    $scope.label = [ '', 'default', 'info', 'primary', 'warning', 'success', 'success', 'danger' ];

    $scope.sentOrders = $scope.getOrders();

    $scope.clearOrders = function(){
      localStorageService.remove('orders')
    };        
    
}
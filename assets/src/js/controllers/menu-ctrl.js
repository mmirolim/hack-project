/**
 * Master Controller
 */
angular.module('Rest')
    .controller('MenuController', ['$scope', 'localStorageService', 'STATUS', 'Menu', '$routeParams', MenuController]);

function MenuController( $scope, localStorageService, $STATUS, Menu, $routeParams ) {

    $scope.menu = $scope.getByCategory( $routeParams.category );    

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

    $scope.addOrder = function( meal ) {        

        var orders = localStorageService.get('orders') || [];

        if( orders.length )
        {
        	orders = _.map( orders, function( obj ){
        	var found = false;
        		if( obj.id == meal.id && obj.serving == meal.serving )
        		{
        			found = true;
        			obj.quantity += meal.quantity; 
        		}

        		return obj 
        	});

        	if( !found ){
        		meal.status = $STATUS.ISSUED;
	        	orders.push( meal );
        	}			        	
        }

        localStorageService.set( 'orders', orders );        
    };       
    
}
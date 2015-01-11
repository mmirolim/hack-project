/**
 * Master Controller
 */

angular.module('Rest')
    .controller('MasterController', [ '$scope', 'localStorageService', '$http',  MasterController ] );

function MasterController( $scope, localStorageService, $http ) {              

    $scope.modalCall = false;

    $scope.getTotalCost = function( meal ){ 

        var cost = 0;

        _.each( $scope.orders ,function( obj ){
           cost += (obj.cost * obj.quantity );
        });            

        return cost;
    };
    
    $scope.getOrders = function(){
        return localStorageService.get('orders') || [];            
    };

    $scope.ordersClass = '';

    $scope.orders = [];

    $scope.getTableId = function(){
        return localStorageService.get('tableId');
    };

    $scope.getOrders = function(){
        return localStorageService.get('orders');
    }; 

    $scope.getByCategory = function( catId ){
        return _.where( $scope.items, { catID: +catId });
    };

    $scope.getItems = function(){
        $http.get('/items').
          success(function( data, status, headers, config) {
            $scope.items = _.map( data, function(obj){ obj.quantity = 0; return obj });
          });
    }

    $scope.getItems();

    $scope.percentService = 10;

    $scope.alerts = [];

   
    
}
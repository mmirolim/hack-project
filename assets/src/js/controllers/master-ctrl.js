/**
 * Master Controller
 */

angular.module('Rest')
    .controller('MasterController', [ '$scope', 'localStorageService', 'Menu',  MasterController ] );

function MasterController( $scope, localStorageService, Menu ) {                   

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
        return _.where( $scope.items, { catId: +catId });
    };

    $scope.percentService = 10;

    $scope.items = Menu.getMenu();       

    $scope.alerts = [];

   
    
}
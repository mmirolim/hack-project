angular.module('Rest')
    .controller('AdminItemController', [ '$scope', '$http',  AdminItemController ] );

function AdminItemController( $scope, $http ) {     

    $scope.item = {};

    $scope.getItems();

    $scope.add = function( item ){
    	$http.post('/items', item ).
	      success( function( data, status, headers, config ) {
	        $scope.getItems();
	      });
    };

    // $scope.del = function( id ){
    // 	$http.delete('/items/' + id ).
	   //    success( function( data, status, headers, config ) {
	   //      getItems();
	   //    });
    // };	    

}
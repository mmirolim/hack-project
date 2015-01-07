angular.module('Rest')
    .controller('AdminCategoryController', [ '$scope', '$http',  AdminCategoryController ] );

function AdminCategoryController( $scope, $http ) {     

   getCategories();

    $scope.add = function( name, desc, staff ){
    	$http.post('/categories', { name: name, desc: desc, staffID: staff }).
	      success( function( data, status, headers, config ) {
	        getCategories();
	      });
    };

    $scope.del = function( id ){
    	$http.delete('/categories/' + id ).
	      success( function( data, status, headers, config ) {
	        getCategories();
	      });
    };	
    
     function getCategories(){
    	$http.get('/categories').
	      success(function(data, status, headers, config) {
	        $scope.categories = data;
	      });
    }

}
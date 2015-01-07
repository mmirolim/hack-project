/**
 * Master Controller
 */

angular.module('Rest')
    .controller('CategoryController', [ '$scope', 'localStorageService', '$http',  CategoryController ] );

function CategoryController( $scope, localStorageService, $http ) {               

    $http.get('/categories').
      success(function(data, status, headers, config) {
        $scope.categories = data;
      });
    
}

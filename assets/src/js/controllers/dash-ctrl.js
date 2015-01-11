/**
 * Master Controller
 */

angular.module('Rest')
    .controller('DashController', [ '$scope', 'localStorageService', 'Category',  DashController ] );

function DashController( $scope, localStorageService, Category ) {               

    $scope.categories = Category.getCategories();
    
}

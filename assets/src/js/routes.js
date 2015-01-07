'use strict';

angular.module('Rest')
    .config( function( $routeProvider, $locationProvider, localStorageServiceProvider ) {
    
    localStorageServiceProvider
            .setPrefix('Rest')
            .setStorageType('sessionStorage')
            .setNotify(true, true); 

    //configuring routes
    $routeProvider
    .when('/menu/:category', {
        templateUrl: 'templates/menu.html',
        controller: 'MenuController',        
    }) 
    //dashboard
    .when('/dash', {
        templateUrl: 'templates/dash.html',
        controller: 'DashController',        
    })    
    //table: where we will read table id
    .when('/table/:tableId', {
        controller: function( $location, $routeParams, tableService, localStorageService ){
            if( tableService.checkTable( $routeParams.tableId ) )
            {   
                localStorageService.set('tableId', $routeParams.tableId );                        

                $location.path("menu");
                
            } else {
                $location.path("/404");
                console.error( $routeParams.tableId, ": Table not found in db" );
            }
        }
    })
    .when('/', {
        templateUrl: 'templates/menu.html',
        controller: 'MenuController'
    })
    .when('/admin/category', {
        templateUrl: 'templates/admin/category.html',
        controller: 'AdminCategoryController'
    });

});
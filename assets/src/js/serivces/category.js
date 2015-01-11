angular.module('Rest')
    .factory('Category', ['$http', CategoryService]);

function CategoryService( $http )
{	
	// function getItems(){
 //    	$http.get('/items').
	//       success(function(data, status, headers, config) {
	//         $scope.items = data;
	//       });
 //    }

	// var getCategories = function() {	
	// 	return [{					
	//         id: 1,
	//         name: "Meals",	        
	//       },{					
	//         id: 2,
	//         name: "Drinks",	        
	//       }];
	// };

	// return {
	// 	getCategories: function() {
	// 		return getCategories();
	// 	}
	// };
};
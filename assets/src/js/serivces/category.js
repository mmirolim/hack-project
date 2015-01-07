angular.module('Rest')
    .factory('Category', ['$http', CategoryService]);

function CategoryService( $http )
{	
	var getCategories = function() {	
		return [{					
	        id: 1,
	        name: "Meals",	        
	      },{					
	        id: 2,
	        name: "Drinks",	        
	      }];
	};

	return {
		getCategories: function() {
			return getCategories();
		}
	};
};
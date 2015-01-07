angular.module('Rest')
    .factory('Menu', ['$http', menuService]);

function menuService( $http )
{	

	var items = [{					
	        id: 1,
	        catId: 1,
	        name: "Palov",
	        desc: "Some desc Some desc Some desc Some desc Some desc Some Sdesc Some desc Some desc Some desc ",
	        serving: 1,
	        image: "img/palov.jpg",
	        cost: 5000,
	        quantity: 0
	      },{
	        id: 6,
	        catId: 1,
	        name: "Palov",
	        desc: "Some desc Some desc Some desc Some desc Some desc Some Sdesc Some desc Some desc Some desc ",
	        serving: 1,
	        image: "img/palov.jpg",
	        cost: 5000,
	        quantity: 0
	      },{
	        id: 2,
	        catId: 1,
	        name: "Kebab",
	        desc: "Some desc",
	        serving: 1,
	        image: "img/kebab.jpg",
	        cost: 3000,
	        quantity: 0
	      },{
	        id: 3,
	        catId: 1,
	        name: "Manti",
	        desc: "Some desc",
	        serving: 1,
	        image: "img/manti.jpg",
	        cost: 3000,
	        quantity: 0
	      },{
	        id: 4,
	        catId: 2,
	        name: "Norin",
	        desc: "Some desc",
	        serving: 1,
	        image: "img/norin.jpg",
	        cost: 3000,
	        quantity: 0
	      },{
	        id: 5,
	        catId: 2,
	        name: "Somsa",
	        desc: "Some desc",
	        serving: 1,
	        image: "img/somsa.jpg",
	        cost: 3000,
	        quantity: 1
	    }];



	var getItems = function() {		
		return items;
	};

	var getByCategory = function( catId ){	
		return _.where(items, { catId: +catId });
	};

	return {
		getMenu: function() {
			return getItems();
		},
		getByCategory: function( catId ){
			return getByCategory( catId );
		},
	};
};
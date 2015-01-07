angular.module('Rest')
    .factory('Orders', ['$http', OrdersService]);

function OrdersService( $http )
{	
	var getOrders = function() {	
		return [{			
	        id: 1,
	        items: [{					
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
		      }],
	        tableId: 1,
	        cost: 8000,
	        percentService: 10,
	        status: 1,	              
	      },{					
	        id: 2,
	        items: [{
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
		      }],
	        tableId: 2,
	        cost: 8000,
	        percentService: 10,
	        status: 1,        
	      }];
	};

	return {
		getOrders: function() {
			return getCategories();
		}
	};
};
angular.module('Rest')
    .factory('tableService', ['$http', tableService]);

function tableService( $http )
{	
	var tablesUrl = '../mock/tables.json';	

	var checkTable = function( tableId ) {
		
		// return $http({
		// 	method: 'JSON',
		// 	url: tablesUrl			
		// });
		return true;
	};

	return {
		checkTable: function( tableId ) {
			return checkTable( tableId );
		}
	};
};
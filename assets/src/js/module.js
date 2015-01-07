angular.module('Rest.config',[])
    .constant('ROLE', { 
    	'ADMIN' : 1, 
    	'MANAGER' : 2, 
    	"STAFF" : 3, 
    	"CLIENT" : 4 } 
    )	
	.constant('STATUS', {
		'ISSUED' : 1, 
		'ACCEPTED' : 2, 
		'INPROGRESS' : 3,
		'READY' : 4,
		'DELIVERED' : 5,
		'PAID' : 6,
		'CANCELLED' : 7} 
	);

angular.module('Rest', ['ngRoute', 'LocalStorageModule','mobile-angular-ui', 'Rest.config']);
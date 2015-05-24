var app = angular.module('myApp', []);
app.controller('MonteController', ['$http', '$log', '$scope', function($http, $log, $scope) {
	$scope.log=$log;
	$log.log('kenny is evil');
	$scope.value = 0,
	$scope.times = 1;
	$scope.myData = {};
	$scope.myData.doClick = function(item, event){
		var responsePromise = $http({
			method: "GET",
			url: "/monte/" + $scope.times			
		});
		responsePromise.success(function(data, status, headers, config) {
                    $scope.value = data.Pi;
                });
                responsePromise.error(function(data, status, headers, config) {
                    alert("AJAX failed!");
                });
	};
}]);

app.directive('validNumber', function() {
  return {
    require: '?ngModel',
    link: function(scope, element, attrs, ngModelCtrl) {
      if(!ngModelCtrl) {
        return; 
      }
	  
      ngModelCtrl.$parsers.push(function(val) {
        if (angular.isUndefined(val)) {
            var val = '';
        }
        var clean = val.replace( /[^0-9]+/g, '');
        if (val !== clean) {
          ngModelCtrl.$setViewValue(clean);
          ngModelCtrl.$render();
        }
        return clean;
      });

      element.bind('keypress', function(event) {
        if(event.keyCode === 32) {
          event.preventDefault();
        }
      });
    }
  };
});
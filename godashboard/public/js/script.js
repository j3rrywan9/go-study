$(document).ready(function () {
	$.ajax({
		type: 'GET',
		url: 'app/list_platform',
		dataType: 'html',
		success: function(data) {
			$("#ajax_output").append(data);
		},
		error: function(xhr, textStatus, errorThrown) {
			$("#ajax_output").append("<h1>" + errorThrown + "</h1>");
		}
	});
})

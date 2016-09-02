$(function () {
  	$('.navbar-toggle-sidebar').click(function () {
  		$('.navbar-nav').toggleClass('slide-in');
  		$('.side-body').toggleClass('body-slide-in');
  		$('#search').removeClass('in').addClass('collapse').slideUp(200);
  	});

  	$('#search-trigger').click(function () {
  		$('.navbar-nav').removeClass('slide-in');
  		$('.side-body').removeClass('body-slide-in');
  		$('.search-input').focus();
  	});
  });
  
  
/*
Please consider that the JS part isn't production ready at all, I just code it to show the concept of merging filters and titles together !
*/
$(document).ready(function(){
    if (document.getElementById("idsearch") && document.getElementById("firstsearch") && document.getElementById("lastsearch") && document.getElementById("emailsearch")){
        document.getElementById("idsearch").style.visibility = "hidden";
        document.getElementById("firstsearch").style.visibility = "hidden";
        document.getElementById("lastsearch").style.visibility = "hidden";
        document.getElementById("emailsearch").style.visibility = "hidden";
        $('.filterable .btn-filter').click(function(){
            var $panel = $(this).parents('.filterable'),
            $filters = $panel.find('.filters input'),
            $tbody = $panel.find('.table tbody');
            if ($filters.prop('disabled') == true) {
                $filters.prop('disabled', false);
                document.getElementById("idsearch").style.visibility = "visible";
                document.getElementById("firstsearch").style.visibility = "visible";
                document.getElementById("lastsearch").style.visibility = "visible";
                document.getElementById("emailsearch").style.visibility = "visible";
                $filters.first().focus();
            } else {
                $filters.val('').prop('disabled', true);
                document.getElementById("idsearch").style.visibility = "hidden";
                document.getElementById("firstsearch").style.visibility = "hidden";
                document.getElementById("lastsearch").style.visibility = "hidden";
                document.getElementById("emailsearch").style.visibility = "hidden";
                $tbody.find('.no-result').remove();
                $tbody.find('tr').show();
            }
        });  
    } else {
        if (document.getElementById("descrsearch") && document.getElementById("pisearch")){
        document.getElementById("descrsearch").style.visibility = "hidden";
        document.getElementById("pisearch").style.visibility = "hidden";
        $('.filterable .btn-filter').click(function(){
            var $panel = $(this).parents('.filterable'),
            $filters = $panel.find('.filters input'),
            $tbody = $panel.find('.table tbody');
            if ($filters.prop('disabled') == true) {
                $filters.prop('disabled', false);
                document.getElementById("descrsearch").style.visibility = "visible";
                document.getElementById("pisearch").style.visibility = "visible";
                $filters.first().focus();
            } else {
                $filters.val('').prop('disabled', true);
                document.getElementById("descrsearch").style.visibility = "hidden";
                document.getElementById("pisearch").style.visibility = "hidden";
                $tbody.find('tr').show();
            }
        });  
        } else if (document.getElementById("ncontrsearch") && document.getElementById("datasearch") && document.getElementById("fornsearch")){
            document.getElementById("ncontrsearch").style.visibility = "hidden";
            document.getElementById("datasearch").style.visibility = "hidden";
            document.getElementById("fornsearch").style.visibility = "hidden";
            $('.filterable .btn-filter').click(function(){
                var $panel = $(this).parents('.filterable'),
                $filters = $panel.find('.filters input'),
                $tbody = $panel.find('.table tbody');
                if ($filters.prop('disabled') == true) {
                    $filters.prop('disabled', false);
                    document.getElementById("ncontrsearch").style.visibility = "visible";
                    document.getElementById("datasearch").style.visibility = "visible";
                    document.getElementById("fornsearch").style.visibility = "visible";
                    $filters.first().focus();
                } else {
                    $filters.val('').prop('disabled', true);
                    document.getElementById("ncontrsearch").style.visibility = "hidden";
                    document.getElementById("datasearch").style.visibility = "hidden";
                    document.getElementById("fornsearch").style.visibility = "hidden";
                    $tbody.find('tr').show();
                }
            });  
        }
    }


    $('.filterable .filters input').keyup(function(e){
        /* Ignore tab key */
        var code = e.keyCode || e.which;
        if (code == '9') return;
        /* Useful DOM data and selectors */
        var $input = $(this),
        inputContent = $input.val().toLowerCase(),
        $panel = $input.parents('.filterable'),
        column = $panel.find('.filters th').index($input.parents('th')),
        $table = $panel.find('.table'),
        $rows = $table.find('tbody tr');
        /* Dirtiest filter function ever ;) */
        var $filteredRows = $rows.filter(function(){
            var value = $(this).find('td').eq(column).text().toLowerCase();
            return value.indexOf(inputContent) === -1;
        });
        /* Clean previous no-result if exist */
        $table.find('tbody .no-result').remove();
        /* Show all rows, hide filtered ones (never do that outside of a demo ! xD) */
        $rows.show();
        $filteredRows.hide();
        /* Prepend no-result row if all rows are filtered */
        if ($filteredRows.length === $rows.length) {
            $table.find('tbody').prepend($('<tr class="no-result text-center"><td colspan="'+ $table.find('.filters th').length +'">No result found</td></tr>'));
        }
    });
});
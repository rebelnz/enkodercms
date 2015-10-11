(function($) {
    'use strict';
    var AppSite = {
        init: function() {
	        AppSite.genericSubmit();
            AppSite.showResetForm();
        },
        genericSubmit: function() {
            $('#form-contact, #form-reset, #form-subscribe').submit(function(e) {
                var process = this.dataset.process;
                // $("#form-" + process).find('input, textarea, button, select').attr("disabled", true);
                $("#" + process + "-confirm").fadeIn().delay(2000).fadeOut();
                $.ajax({
                    url: '/' + process,
                    dataType: 'json',
                    type: 'post',
                    contentType: 'application/x-www-form-urlencoded',
                    data: $(this).serialize(),
                    success: function(data){
                        $("#form-" + process)[0].reset();
                        // $("#form-" + process).find('input, textarea, button, select').attr("disabled", false);
                    }
                });
                e.preventDefault();                
            });
        },    
        showResetForm: function() {
            $('#reset-show').click(function() {
                $('#form-reset').slideToggle();
            });
        }
    }
    $(function() {
        AppSite.init();
    });
})(jQuery);

(function($) {
    'use strict';
    var AppSite = {
        init: function() {
	        AppSite.genericSubmit();
	        AppSite.showResetForm();
        },
        // use data-process to get str "contact" or "reset"
        genericSubmit: function() {
            $('#form-contact, #form-reset').submit(function(e) {
                var process = this.dataset.process;
                $.ajax({
                    url: '/' + process,
                    dataType: 'json',
                    type: 'post',
                    contentType: 'application/x-www-form-urlencoded',
                    data: $(this).serialize(),
                    success: function(data){
                        $("#" + process + "-confirm").fadeIn().delay(2000).fadeOut();
                        console.log(data);
                        $("#form-" + process)[0].reset();
                    }
                });
                e.preventDefault();                
            });
        },
        showResetForm: function() {
            $('#reset-show').click(function() {
                $('#form-reset').fadeToggle();
            });
        }
    }
    $(function() {
        AppSite.init();
    });
})(jQuery);

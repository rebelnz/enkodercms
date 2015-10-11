(function($) {
    'use strict';
    var AppAdmin = {        
        init: function() { // add functions here
	        AppAdmin.preventScroll();
	        AppAdmin.toggleProfile();
	        AppAdmin.deleteConfirm();            
	        AppAdmin.sortNavigation();            
	        AppAdmin.customTrumbowyg();            
	        AppAdmin.updateStatus();        
	        AppAdmin.updateSettings();            
        },
        
        preventScroll: function() {
            if (location.hash) {
                setTimeout(function() {
                    window.scrollTo(0, 0);
                }, 0);
            }        
        },

        toggleProfile: function() {
            $("#nav-profile").on("click",function(e) {
                e.preventDefault();
                $("#admin-profile-data").fadeToggle();
            });
        },
        
        deleteConfirm: function() {
            $('.delete-item').on('click', function () {
                return confirm('Are you sure you want to delete this item? This cannot be undone.');
            });
        },
        
        sortNavigation: function() {
            $(".sortable").sortable({
                connectWith: ".sortable",
                opacity: 0.5,
                update: function(event, ui) {
                    if (this === ui.item.parent()[0]) {
                        var order = $(this).sortable('toArray');
                        var type = ui.item.parent().attr('id');
                        $.post('/admin/ajax/navorder',{'navorder[]':order,'navtype':type});
                    }
                }
            });            
        },

        customTrumbowyg: function() { // wysiwyg editor
            $('#content').trumbowyg({
                resetCss: true,
                autogrow: true
            });
        },

        updateStatus: function() { // update page - live/draft
            $('.update-status').on("click",function(e) {
                e.preventDefault();
                var link = $(this);
                var url = link.attr("href");
                var span = link.parent();
                $.get(url)
                    .done(function() {
                        span.toggleClass('status-live status-draft');
                        link.html(link.text() == 'Live' ? 'Draft' : 'Live');
                    });
                
            });
        },

        
        updateSettings: function() {
            var viewModel = {};
            var data = $.getJSON( "/admin/ajax/getsettings", function( data ) {
                viewModel = {
                    email: ko.observable(data.Email),
                    contact: ko.observable(data.Contact),
                    address: ko.observable(data.Address),
                    street: ko.observable(data.Street),
                    suburb: ko.observable(data.Suburb),
                    city: ko.observable(data.City),
                    code: ko.observable(data.Code),
                    facebook: ko.observable(data.Facebook),
                    twitter: ko.observable(data.Twitter),
                    linkedin: ko.observable(data.Linkedin),
                    description: ko.observable(data.Description),
                    keywords: ko.observable(data.Keywords),
                    ganalytics: ko.observable(data.Ganalytics),
                    smtp: ko.observable(data.Smtp)
                };
                ko.applyBindings(viewModel);
            });
            $(".tab-content input").prop('disabled', true); // disable inputs
            $("#admin-settings-edit").click(function() {
                $(".admin-settings-input, .admin-settings-display").toggle();        
                var editText = $('#admin-settings-edit').text(); 
                var bgColor = $('#admin-settings-edit').css("background-color");
                $("#admin-settings-edit").text(editText === "Edit" ? "Done" : "Edit") 
                    .css({"background-color":bgColor === "rgb(46, 184, 31)" ? "rgb(17, 165, 224)" : "rgb(46, 184, 31)"});        
                $('.tab-content input').prop('disabled', function(i, v) { return !v; })
                if ( editText === "Done" ) {
                    var jsData = ko.toJS(viewModel);            
                    $.get("/admin/ajax/updatesettings", jsData )
                        .done(function() {
                            $('#settings-updated').show().delay('2000').fadeOut();
                        });
                }
            });    

        },

    }
    
    $(function() {
        AppAdmin.init();
    });    
})(jQuery);

$(function(){
    $.ajaxSetup({
        beforeSend: function() {
            $('#loader').show();
        },
        complete: function(){
            $('#loader').hide();
        },
        success: function() {}
    });    
    getMapData();
});

function buildMap(mapData) {
    var map = L.map('map').setView([mapData['Latitude'],
                                    mapData['Longitude']], 13);
    L.tileLayer('http://otile4.mqcdn.com/tiles/1.0.0/map/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
    var marker = L.marker([ mapData['Latitude'], mapData['Longitude']]).addTo(map);
}

function getMapData() {
    $.get("/admin/ajax/getmap", function(data) {
        buildMap(data);
    },'json');
}

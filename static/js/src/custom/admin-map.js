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

    var map = L.map('map').setView([mapData['Latitude'],mapData['Longitude']], 13);
    
    L.tileLayer('http://otile4.mqcdn.com/tiles/1.0.0/map/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
    
    var marker = L.marker([ mapData['Latitude'], mapData['Longitude']], {draggable:true}).addTo(map);

    marker.on('dragend', function(event) {
        updateMap(marker.getLatLng());
        // var marker = event.target;  // you could also simply access the marker through the closure
        // var result = marker.getLatLng();  // but using the passed event is cleaner
    });
}

function updateMap(latlng) {
    $.post("/admin/ajax/updatemap", { latitude: latlng.lat, longitude: latlng.lng } )
        .done(function() {
            $('#map-updated').show().delay('2000').fadeOut();
        });
}

function getMapData() {
    $.get("/admin/ajax/getmap", function(data) {
        buildMap(data);
    },'json');
}

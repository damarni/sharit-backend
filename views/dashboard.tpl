<!DOCTYPE html>
<html>
  <head>
    <style type="text/css">
      html, body { height: 100%; margin: 0; padding: 0; font-family: Helvetica}
      #map { height: calc(100% - 40px); }
    </style>

    <style>
          html, body {
            height: 100%;
            margin: 0;
            padding: 0;
          }
          #map {
            height: 100%;
          }
    #floating-panel {
      position: absolute;
      top: 10px;
      left: 25%;
      z-index: 5;
      background-color: #fff;
      padding: 5px;
      border: 1px solid #999;
      text-align: center;
      font-family: 'Roboto','sans-serif';
      line-height: 30px;
      padding-left: 10px;
    }

          #floating-panel {
            background-color: #fff;
            border: 1px solid #999;
            left: 25%;
            padding: 5px;
            position: absolute;
            top: 10px;
            z-index: 5;
          }
        </style>

    <script src="https://code.jquery.com/jquery-2.2.3.min.js"></script>
  </head>
  <body>
    <div style="width: 100%; height: 20px; background: #333; color: white; padding: 10px; line-height: 20px">
      <span>Sharit Dashboard</span>

    </div>
    <div id="floating-panel">
      <button onclick="toggleHeatmap()">Toggle Heatmap</button>
      <button onclick="changeGradient()">Change gradient</button>
      <button onclick="changeRadius()">Change radius</button>
      <button onclick="changeOpacity()">Change opacity</button>
    </div>
    <div id="map"></div>

    <script type="text/javascript">


    var map, heatmap;

    function initMap() {
      map = new google.maps.Map(document.getElementById('map'), {
        zoom: 14,
        center: {lat: 41.386756, lng: 2.170467},
        mapTypeId: google.maps.MapTypeId.SATELLITE,
        setMap : map

      });

      heatmap = new google.maps.visualization.HeatmapLayer({
        data: getPoints(),
        map: map
      });
    }

    function toggleHeatmap() {
      heatmap.setMap(heatmap.getMap() ? null : map);
    }

    function changeGradient() {
      var gradient = [
        'rgba(0, 255, 255, 0)',
        'rgba(0, 255, 255, 1)',
        'rgba(0, 191, 255, 1)',
        'rgba(0, 127, 255, 1)',
        'rgba(0, 63, 255, 1)',
        'rgba(0, 0, 255, 1)',
        'rgba(0, 0, 223, 1)',
        'rgba(0, 0, 191, 1)',
        'rgba(0, 0, 159, 1)',
        'rgba(0, 0, 127, 1)',
        'rgba(63, 0, 91, 1)',
        'rgba(127, 0, 63, 1)',
        'rgba(191, 0, 31, 1)',
        'rgba(255, 0, 0, 1)'
      ]
      heatmap.set('gradient', heatmap.get('gradient') ? null : gradient);
    }

    function changeRadius() {
      heatmap.set('radius', heatmap.get('radius') ? null : 20);
    }

    function changeOpacity() {
      heatmap.set('opacity', heatmap.get('opacity') ? null : 0.2);
    }

    function getPoints() {
      return [      {{ range .data }}
             new google.maps.LatLng({{.Lat}},{{.Lng}}),
             {{ end }}]
    }





</script>

<script async defer
    src="https://maps.googleapis.com/maps/api/js?key=AIzaSyDb3m6o_OKV5OwUsLAdzdDQB8DQ6BJMIl0&callback=initMap&libraries=visualization">
</script>
</body>
</html>

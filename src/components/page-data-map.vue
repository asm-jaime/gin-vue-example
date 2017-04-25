<template>
  <div>
    <svg id="map" class="map">
    </svg>
  </div>
</template>

<script>

  import { mapGetters, mapActions } from 'vuex'
import * as gets from '../constants/types.getters.js'
import * as acts from '../constants/types.actions.js'

export default {
  name: 'PageDataMap',
  data(){ // {{{
    return {
      cities: [
        { city: "ZANZIBAR", code: "ZNZ", country: "TANZANIA", lat: "-6.13", lon: "39.31" },
        { city: "ZANZ", code: "ZNZ", country: "TANZANIA", lat: "-7.13", lon: "39.31" },
        { city: "ZANZIR", code: "ZNZ", country: "TANZANIA", lat: "-8.13", lon: "19.31" },
        { city: "ZANZR", code: "ZNZ", country: "TANZANIA", lat: "6.13", lon: "16.31" },
        { city: "ZANZAR", code: "ZNZ", country: "TANZANIA", lat: "8.13", lon: "19.31" },
      ],
    }
  }, // }}}
  mounted: function() { // {{{
    console.log('##we mounted');
    this.GET_DATA().then(()=>{
      this.draw();
    })
    // window.addEventListener('resize', this.handleResize);
  }, // }}}
  computed: { // {{{
    ...mapGetters([
      gets.DATA,
    ]),
  }, // }}}
  methods: {
    ...mapActions([ // {{{
      acts.GET_DATA,
    ]), // }}}
    handleResize: function() { // {{{
      //this.draw();
    }, // }}}
      draw: function() { // {{{
        const d3 = this.$d3;
        const height = window.outerHeight;
        const width = window.outerWidth;

        const projection = d3.geoMercator()
          .translate([0, 0])
          .scale(width / 2 / Math.PI);

        const zoom = d3.zoom()
          .scaleExtent([1, 12])
          .on("zoom", move);

        const path = d3.geoPath().projection(projection);

        const svg = d3.select("#map")
          .attr("width", width)
          .attr("height", height)
          .call(zoom);

        const g = svg.append("g");

        d3.json("world-110m.json", (error, topology) => {
          g.selectAll("path")
            .data(topojson.feature(topology, topology.objects.countries).features)
            .enter()
            .append("path")
            .attr("class", "land")
            .attr("d", path);

          g.selectAll("circle") // {{{
            .data(this.DATA)
            .enter()
            .append("a")
            .attr("xlink:href", (d) => "https://www.google.com/search?q=" + d.data )
            .append("circle")
            .attr("cx",(d) => projection([d.location.coordinates[0], d.location.coordinates[1]])[0])
            .attr("cy",(d) => projection([d.location.coordinates[0], d.location.coordinates[1]])[1])
            .attr("r", 5)
            .style("fill", "red"); // }}}

        });
        function move() {
            const t = d3.event.transform;
            //t.x = Math.min(width/2*(t.k-1), Math.max(width/2*(1-t.k), t.x));
            //t.y = Math.min(height/2*(t.k-1), Math.max(height/2*(1-t.k), t.y));
            //t.y = Math.min(height/2*(t.k-1)+50*t.k, Math.max(height/2*(1-t.k)-150*t.k, t.y));
            t.x = Math.max(0, Math.min(t.x, width - t.k*50));
            t.y = Math.max(0, Math.min(t.y, height - t.k*50));
            console.log(t);
            g.attr("transform", "translate("+t.x+','+t.y+")scale("+t.k+")");
            g.selectAll("path").attr("d", path.projection(projection));
            g.selectAll("circle").attr("d", path.projection(projection));
        };
      }, // }}}
    showConstants: function(){
      console.log("## some msg");
    }
  }
}

</script>

<style>

  .overlay {
    fill: none;
    pointer-events: all;
  }

  .boundary {
    fill: none;
    stroke: #fff;
    stroke-linejoin: round;
    stroke-linecap: round;
  }

  .land {
    fill: #000;
  }

  path {
    stroke: white;
    stroke-width: 0.25px;
    fill: grey;
  }

  .map{
    z-index: 0;
    position: absolute;
  }
</style>

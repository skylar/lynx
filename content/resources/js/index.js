app = new Vue({
  el: "#app",
  data: {
    bait: '',
  },
  methods: {
    hunt: function() {
      if (this.bait == "") {
        return;
      }
      window.location = "/s/" + encodeURIComponent(this.bait);
    },
  },
});

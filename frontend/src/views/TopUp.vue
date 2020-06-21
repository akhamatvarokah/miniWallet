<template>
  <div class="container">
    <header class="jumbotron">
      <h3>Top Up Balance</h3>
      <form>
        <div class="form-group">
          <label for="exampleInputEmail1">Amount</label>
          <input name="amount" v-model="amount" type="number" class="form-control" id="exampleInputEmail1" aria-describedby="amount" placeholder="Enter Amount">
          <small id="amount" class="form-text text-muted">Please Input Amount</small>
        </div>
        <button v-on:click="sendAmount" type="button" class="btn btn-primary">Submit</button>
      </form>
    </header>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Top Up',
  data() {
    return {
      results: ''
    };
  },
  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    }
  },
   methods: {
    sendAmount: function () {
      // `this` inside methods points to the Vue instance
      let amount = parseInt(document.querySelector("input[name=amount]").value);
      const params = {
            amount: amount
        };

      if (amount > 0) {
        var config = {
        headers: { 
          'Content-Type': 'application/x-www-form-urlencoded',
          'Authorization': 'Bearer ' + this.currentUser.token
        
        }
      };

      axios.post(this.$baseUrl + 'balance/topup', params,config).then(response => {
        if (response.status == "200") {
          alert("sukses top up");
          location.href = "home";
        }
      })
      }
    }
  },
  mounted() {
    if (!this.currentUser) {
      this.$router.push('/login');
    }
  }
};
</script>

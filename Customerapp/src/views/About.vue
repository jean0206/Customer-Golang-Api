<template>
  <div class="about">
    <div class="list-products">
      Productos comprados por {{ customer.name }}

      <Carousel v-if="products.length > 4" :products="products"></Carousel>
    </div>

    <div class="list-products">
      Productos que te pueden interesar

      <Carousel
        v-if="suggestions.length > 1"
        :products="suggestions"
      ></Carousel>
    </div>

    <div class="customer-ip">
      <div class="customer-title">
        Usuarios que compraron desde la misma IP
      </div>
      <div class="customer-ip-content">
        <div class="customer-card" v-for="(item, key) in buyers" :key="key">
          <CardUser
            :name="item.name"
            :age="item.age"
            :id="item.id"
            type="suggestion"
          ></CardUser>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import Carousel from "../components/Carousel";
import CardUser from "../components/CardUser.vue";

export default {
  components: {
    Carousel,
    CardUser,
  },
  data: function() {
    return {
      newProduct: {
        nameProduct: "",
        idProduct: "",
        price: 0,
      },
      customer: {},
      transactions: [],
      idProducts: [],
      products: [],
      buyers: [],
      suggestions: [],
    };
  },
  async created() {
    this.customer = JSON.parse(localStorage.getItem("customerInfo"));
    await axios
      .get("http://localhost:3000/transaction/" + this.customer.id)
      .then((response) => {
        this.transactions = response.data["customer"];
      })
      .catch((error) => {
        console.error(error.message);
      });
    // console.log(this.transactions)
    this.transactions.forEach((transaction) => {
      transaction.products.forEach((product) => {
        this.idProducts.push(product);
      });
    });
    this.idProducts.forEach(async (id) => {
      await axios
        .get("http://localhost:3000/product/" + id)
        .then((response) => {
          this.newProduct = {
            nameProduct: response.data["product"][0].nameProduct,
            idProduct: response.data["product"][0].idProduct,
            price: response.data["product"][0].price,
          };

          this.products.push(this.newProduct);
        });
    });

    await axios
      .get("http://localhost:3000/customer/ip/" + this.transactions[0].ip)
      .then((response) => {
        response.data["customer"].map((customer) => {
          axios
            .get("http://localhost:3000/customer/" + customer.buyerId)
            .then((response) => {
              this.buyers.push(response.data["customer"][0]);
            });
        });
      });
    console.log(this.buyers);

    const prices = await this.products.map((product) => product.price);
    const min = Math.min.apply(null, prices);
    const max = Math.max.apply(null, prices);
    await axios
      .get("http://localhost:3000/product/suggestion/" + min + "/" + max)
      .then((response) => {
        this.suggestions = response.data["products"];
      });
  },
};
</script>

<style scoped>
.about {
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;
}
.list-products {
  font-size: 40px;
  margin-top: 80px;
}
.customer-card {
  margin-top: 30px;
  margin-left: 20px;
}

.customer-ip-content {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.customer-title {
  font-size: 40px;
}

.customer-ip {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
  width: 100%;
  justify-content: center;
  margin-top: 50px;
}
</style>

<template>
  <div class="carousel">
    <div class="carousel-btn">
      <button>
        <font-awesome-icon
          icon="angle-left"
          :style="{ color: '#424AF5' }"
          size="2x"
          @click="moveLeft"
        />
      </button>
    </div>
    <div  class="carousel-content">

             <div      
        v-for="(item, key) in selectProducts"
        :key="key"
        class="card-product"
      >
        <div class="card-image">
            <img src="../assets/product.webp"/>
        </div>

        <div class="card-name">
            {{item.nameProduct}}
        </div>
        <div class="card-price">
           $ {{item.price}}
        </div>


        </div>
     
    </div>
    <div class="carousel-btn">
      <button id="button" ref="myBtn">
        <font-awesome-icon
          icon="angle-right"
          :style="{ color: '#424AF5' }"
          size="2x"
          @click="moveRight"
          
        />
      </button>
    </div>
  </div>
</template>

<script>
export default {
props:{
    products:Array
},
  data: () => {
    return {     
      actual: 0,
      selectProducts:[],
      amount: 5,
      width: window.innerWidth,
    };
  },
   beforeCreate() {
          window.onresize = () => {
        this.resize()
      }

   },
  created(){
        this.addData()
      this.resize()
  },
  mounted() {
   this.width = window.innerWidth;
      
  },
  methods: {
  
      resize(){
          this.width = window.innerWidth;
      if (this.width < 1356 && this.width > 1053) {
        this.amount = 3;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      } else if (this.width < 1053 &&this.width>730) {
        this.amount = 2;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      } else if(this.width<730) {
          this.amount = 1;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      }
      else if (this.width > 1356) {
        this.amount = 4;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      }else {
           this.amount = 4;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      }
      },
    moveRight() {

      if (this.actual + this.amount < this.products.length) {
        this.actual += this.amount;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      }else{
          this.selectProducts = this.products.slice(
          this.actual,
          this.products.length
        );
      }
    },
    addData(){
        console.log(this.products.length)
        this.selectProducts= this.products.slice(0,this.amount);
        this.resize()
    },
    moveLeft() {

        console.log(this.actual)
      if (this.actual - this.amount >= 0) {
        this.actual -= this.amount;
        this.selectProducts = this.products.slice(
          this.actual,
          this.actual + this.amount
        );
      }else{
          this.actual=0
           this.selectProducts = this.products.slice(
          0,
           this.amount)
      }
    },
  },
};
</script>

<style scoped>
.carousel {
  width: 100%;
  height: 310px;
  display: flex;
  justify-content: space-between;
  background-color: white;
}
.carousel-btn button {
  width: 60px;
  height: 310px;
  outline: none;
  border: none;
  background-color: transparent;
  cursor: pointer;
}
.carousel-btn :hover {
  background-color: rgba(216, 191, 216, 0.082);
}
.card-product {
  width: 300px;
  height: 300px;
  background-color: rgba(216, 191, 216, 0.281);
  border-radius: 18px;
  margin-left: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-around;
}
.carousel-content {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-name{
    font-size: 20px;
    
    text-align: center;
    justify-content: center;
    font-weight: lighter;
}

.card-price{
    margin-top: 10px;
    font-size: 26px;
}

.card-image{
    overflow: hidden;
}
.card-image img{
    width: 100px;
    height: 100px;
    border-radius: 500px;
}

</style>

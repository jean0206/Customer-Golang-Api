<template>
<div class="card">
  <div class="card-container">
    <div class="card-image">
      <img src="../assets/avatar.png"/>
    </div>
    <div class="card-content">
      <div>ID: {{ id }}</div>
      <br />
      <div>Nombre: {{ name }}</div>
      <br />
      <div>Edad: {{ age }}</div>
      <br />
    </div>
    <div v-if="type!='suggestion'" class="card-action">
      <button @click="clicked">
        <font-awesome-icon
          icon="angle-right"
          :style="{ color: '#424AF5' }"
          size="2x"
        />
      </button>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: "CardUser",
  props: {
    name: String,
    age: Number,
    id: String,
    type:String
  },
  methods: {
      clicked(){
          this.$emit('clicked')
      }
  },
};
</script>

<style scoped>
.card-container {
  
  border-radius: 16px;
  background-color: #fff;
  box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.04), 0px 2px 6px rgba(0, 0, 0, 0.04),
    0px 0px 1px rgba(0, 0, 0, 0.04);
  display: flex;
  
}

.card-container :hover {
 

}



.card-image {
  width: 136px;
  height: 136px;
  left: -12px;
  top: 0px;
  border-radius: 16px;
  overflow: hidden;
}

.card-image img{
  width: 136px;
  height: 136px;
  left: -12px;
}
.card-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  text-align: center;
  width: 207px;
  padding-left: 10px;
}

.card-action {
  display: flex;
  border: none;
  outline: none;
}

.arrow {
  height: 100px;
}

.card-action button {
  display: flex;
  border: none;
  outline: none;
  text-align: center;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  background-color: transparent;
}
</style>

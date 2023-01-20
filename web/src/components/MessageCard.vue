<template>
  <div class="card-wrapper">
    <div class="card-icon"><img src="../assets/img/message.png" alt="logo" /></div>
    <div class="card-main">
      <div class="time">
        <span>{{ time }}</span>
      </div>
      <div class="title">
        <span>{{ title }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="js">
export default {
  name: "Card",
  props: {
    time: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
  },
  setup() {
    return {};
  },
};
</script>

<style scoped>
.card-wrapper {
  width: 100%;
  height: 10vh;
  padding: 10px;
  display: flex;
  border-bottom: 1px solid #eeeeee;
}
.card-wrapper .card-icon {
  width: 20%;
  height: 100%;
  object-fit: cover;
}
.card-wrapper .card-icon img {
  height: 100%;
}
.card-wrapper .card-main {
  flex: 1;
  display: flex;
  flex-direction: column;
}
.card-wrapper .card-main .time,
.card-wrapper .card-main .title {
  height: 50%;
}
.card-wrapper .card-main .time {
  color: #aaaaaa;
  font-size: 0.8em;
  display: flex;
  align-items: flex-end;
}
.card-wrapper .card-main .title {
  margin-top: 5px;
  font-weight: 600;
}
</style>

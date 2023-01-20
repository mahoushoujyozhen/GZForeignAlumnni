<template>
  <div class="card-wrapper">
    <div class="card-main">
      <div class="top">
        <div class="title">
          <span>{{ title }}</span>
        </div>
        <div class="time">
          <span>{{ time }}</span>
        </div>
      </div>
      <div class="bottom">
        <span>{{ content }}</span>
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
    content: {
      type: String,
      required: true
    }
  },
  setup() {
    return {};
  },
};
</script>

<style scoped>
@media screen and (min-width: 600px) {
  .card-wrapper {
    width: 100%;
    padding: 20px;
    display: flex;
    align-items: center;
  }
  .card-wrapper .card-main .top {
    display: flex;
  }
  .card-wrapper .card-main .top .time,
  .card-wrapper .card-main .top .title {
    font-size: 16px;
  }
  .card-wrapper .card-main .top .time {
    color: #999;
    margin: 0 15px;
  }
  .card-wrapper .card-main .title {
    font-weight: 600;
  }
  .card-wrapper .card-main .bottom {
    padding-top: 8px;
    color: #666;
    font-size: 16px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
  }
}
@media screen and (max-width: 600px) {
  .card-wrapper {
    width: 100%;
    height: 10vh;
    padding: 10px;
    display: flex;
    align-items: center;
    border-radius: 5px;
    background-color: white;
  }
  .card-wrapper .card-icon img {
    height: 90%;
  }
  .card-wrapper .card-main {
    flex: 1;
    display: flex;
    flex-direction: column;
  }
  .card-wrapper .card-main .top .time,
  .card-wrapper .card-main .top .title {
    height: 50%;
  }
  .card-wrapper .card-main .time {
    font-size: 0.35rem;
    color: #aaaaaa;
    display: flex;
    align-items: flex-end;
  }
  .card-wrapper .card-main .top .title {
    font-size: 0.45rem;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
  }
  .card-wrapper .card-main .bottom {
    display: none;
  }
  .title{
    line-height: 2;
  }
}
</style>

<template>
  <div class="mobileNav"><span>{{ title }}</span></div>
</template>

<script>
export default {
  name: "MobileNav",
  props:['title']
}
</script>

<style lang="less" scoped>
@baseFont:37.5;
@media screen and (max-width: 600px) {
  .mobileNav{
    position: relative;
    height: (40rem/@baseFont);
    line-height: (40rem/@baseFont);
    background-color: white;
    border-bottom: 1px solid #efefef;
    font-size: (18rem/@baseFont);
    text-align: center;
  }
}
@media screen and (min-width: 600px){
  .mobileNav{
    display: none;
  }
}
</style>

<template>
<div class="signin">
  <div class="ball"></div>

  <Rocket></Rocket>

  <div class="slogan">
    <div class="text" v-t="'SIGNIN.SLOGAN'"></div>
  </div>

  <div class="panel">
    <div class="form">
      <div class="header" v-t="'SIGNIN.SIGNIN'"></div>
      <div class="user-name">
        <YuumiInput size="lg" round v-model="username" :placeholder="$t('SIGNIN.USER_NAME')">
          <template #prefixIcon>
            <YuumiIcon icon="line-user" style="color: #aaa;font-size: 16px;"></YuumiIcon>
          </template>
        </YuumiInput>
      </div>

      <div class="password">
        <YuumiInput type="password" size="lg" round v-model="password" :placeholder="$t('SIGNIN.PASSWORD')" @keyup.enter="onSignin">
          <template #prefixIcon>
            <YuumiIcon icon="line-lock" style="color: #aaa;font-size: 16px;"></YuumiIcon>
          </template>
        </YuumiInput>
      </div>

      <div class="submit">
        <YuumiButton size="lg" round theme="primary"
          @click="onSignin"
        >{{signinBtnText}}</YuumiButton>
      </div>
    </div>
  </div>
</div>
</template>

<script lang="ts" setup>
import Rocket from "@/components/rocket/index.vue"
import { SigninAPI } from "@/api"
import { useStore } from "@/store"
import { computed, ref } from "vue"
import { useRouter } from "vue-router"
import { useI18n } from "vue-i18n"
import { createMessage } from "yuumi-ui-vue"
const store = useStore()
const router = useRouter()
const { t } = useI18n()

/** 登录 */
const loading = ref(false)
const username = ref("")
const password = ref("")
const signinBtnText = computed(() => {
  return loading.value ? t("SIGNIN.SIGNING") : t("SIGNIN.SIGNIN")
})

function onSignin() {
  if (!username.value || !password.value) {
    createMessage({
      theme: "warn",
      message: t("SIGNIN.FROM_VALIDATOR_MSG")
    })
    return
  }

  if (loading.value) return
  loading.value = true

  return SigninAPI.signinWithPassword({
    name: username.value,
    password: password.value
  }).then(({ token, administrator }) => {
    store.dispatch("global/updateToken", token)
    store.dispatch("global/updateUserData", administrator)
    router.replace("/")
  }).finally(() => {
    loading.value = false
  })
}
</script>

<style lang="scss" scoped>
.signin {
  position: relative;
  width: 100vw;
  height: 100vh;

  .slogan {
    position: fixed;
    left: 23%;
    top: 10%;
    transform: translateX(-50%);

    font-size: 28px;
    font-weight: bold;
    color: #ffffff;

    > .text{
      white-space: pre-line;
    }
  }

  .panel {
    width: 50%;
    padding: 24px;

    position: absolute;
    left: 75%;
    top: 50%;
    transform: translate3d(-50%, -50%, 0);
  }
}

.form {
  margin: 0 auto;
  width: 300px;

  .header {
    font-size: 24px;
    text-align: center;
  }

  .header, .user-name, .password {
    margin-bottom: 16px;
  }

  .yuumi-input {
    width: 100%;
    border-radius: 50px;
  }

  .yuumi-button {
    width: 100%;
  }
}

:deep(.rocket) {
  position: fixed;
  bottom: 50%;
  left: 25%;
  transform: scale(0.8) rotate(20deg) translate3d(-50%, 100%, 0);
}

.ball {
  width: 100vw;
  height: 100vh;

  position: fixed;
  left: -100%;
  top: -60%;
  transform: scale(2, 4);
  z-index: -1;

  border-radius: 50%;
  background-color: #0d6efd;
}
</style>

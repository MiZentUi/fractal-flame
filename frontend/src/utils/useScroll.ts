import { readonly, ref } from "vue";

const scrollProgress = ref<number>(0)

document.addEventListener("scroll",  () => {
    const scrollTop = window.scrollY;
    const docHeight = document.documentElement.scrollHeight;
    const winHeight = window.innerHeight;
    scrollProgress.value = scrollTop / (docHeight - winHeight);
});

const useScroll = () => {
    return {
        scrollProgress: readonly(scrollProgress)
    }
}

export {useScroll}

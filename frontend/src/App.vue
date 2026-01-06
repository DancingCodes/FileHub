<template>
    <div class="fullBox">
        <div class="uploadTitle">Powered by FileHub</div>
        <div class="uploadMain" :class="{ 'loadingState': isUploading }" @click="handleClick" @dragover.prevent
            @dragenter.prevent @drop.prevent="handleDrop">
            <input ref="fileIptRef" type="file" class="fileIpt" @change="handleFileChange" />
            <div class="uploadTips">
                {{ isUploading ? `上传进度：${uploadProgress}%` : '点击、拖拽或粘贴文件至此处，即刻获取在线链接' }}
            </div>
        </div>
        <div class="uploadResult" v-if="fileUrl">
            <div class="fileUrlText">{{ fileUrl }}</div>
            <div class="copyBtn" @click="copyText(fileUrl)">复制</div>
        </div>
    </div>

    <Toaster position="top-center" />
</template>
<script setup lang="ts">
import { Toaster, toast } from 'vue-sonner'
import { uploadFileApi } from '@/api/file'
import { ref, onMounted, onUnmounted } from 'vue'
onMounted(() => {
    window.addEventListener('paste', handlePaste)
})
onUnmounted(() => {
    window.removeEventListener('paste', handlePaste)
})


const fileIptRef = ref<HTMLInputElement | null>(null)
const isUploading = ref(false)
const fileUrl = ref('')
const uploadProgress = ref(0)

const handleClick = () => {
    if (isUploading.value) return
    if (!fileIptRef.value) return
    fileIptRef.value.click()
}

const handleDrop = (e: DragEvent) => {
    if (isUploading.value) return
    const files = e.dataTransfer?.files
    if (files?.length) {
        uploadFile(files[0]!)
    }
}

const handlePaste = (e: ClipboardEvent) => {
    if (isUploading.value) return

    const files = e.clipboardData?.files
    if (files?.length) {
        uploadFile(files[0]!)
    }
}


const handleFileChange = async (e: Event) => {
    const target = e.target as HTMLInputElement
    const file = target.files?.[0]
    if (!file) return
    uploadFile(file)
}


const uploadFile = (file: File) => {
    uploadProgress.value = 0
    isUploading.value = true
    const formData = new FormData()
    formData.append('file', file)


    uploadFileApi(formData, (progress) => {
        uploadProgress.value = progress
    }).then((res) => {
        fileUrl.value = `http://filehubapi.moonc.love/uploads/${res.data.file_uuid}`
    }).finally(() => {
        isUploading.value = false
        if (fileIptRef.value) fileIptRef.value.value = ''
    })
}

const copyText = (text: string) => {
    navigator.clipboard.writeText(text)
    toast.success('已复制到剪贴板')
}


</script>
<style scoped lang="scss">
.fullBox {
    background-color: #1a1a1a;
    color: #f0f0f0;
    position: absolute;
    left: 0;
    top: 0;
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    .uploadTitle {
        font-size: 48px;
        font-weight: bold;
        margin-bottom: 30px;
    }

    .uploadMain {
        width: 60%;
        height: 60%;
        border-radius: 16px;
        border: 2px dashed #ccc;
        cursor: pointer;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        .fileIpt {
            display: none;
        }

        @keyframes pulseGlow {
            0% {
                border-color: rgba($color: #ccc, $alpha: 0.5);
                box-shadow: 0 0 40px rgba(255, 0, 122, 0.2);
            }

            50% {
                border-color: rgba($color: #ccc, $alpha: 1);
                box-shadow: 0 0 70px rgba(255, 0, 122, 0.5);
            }

            100% {
                border-color: rgba($color: #ccc, $alpha: 0.5);
                box-shadow: 0 0 40px rgba(255, 0, 122, 0.2);
            }
        }

        &.loadingState {
            animation: pulseGlow 1.5s infinite;
            cursor: wait;
        }
    }


    .uploadResult {
        margin-top: 20px;
        border-radius: 8px;
        border: 2px dashed #ccc;
        width: 60%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        padding: 6px 12px;
        box-sizing: border-box;
        gap: 8px;


        .fileUrlText {
            flex: 1;
            overflow: hidden;
            text-overflow: ellipsis;
            white-space: nowrap;
            font-size: 20px;
        }

        .copyBtn {
            background-color: #ccc;
            color: #1a1a1a;
            border-radius: 4px;
            padding: 2px 12px;
            box-sizing: border-box;
            cursor: pointer;
            transition: all 0.3s;
            font-size: 15px;

            &:hover {
                background-color: #fff;
            }
        }
    }

}
</style>

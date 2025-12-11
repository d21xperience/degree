// src/composables/useDialogStatus.js
import { ref } from 'vue';


const dialogState = ref({
    visible: false,
    type: 'success',
    title: '',
    message: '',
    buttonLabel: 'OK',
    onConfirm: null
});
export function useDialogStatus() {
    const show = ({ type = 'success', title, message = '', buttonLabel = 'OK', onConfirm = () => {} } = {}) => {
        dialogState.value = {
            visible: true,
            type,
            title,
            message,
            buttonLabel,
            onConfirm
        };
    };

    const hide = () => {
        dialogState.value.visible = false;
    };

    const confirm = () => {
        dialogState.value.onConfirm?.();
        hide();
    };

    return {
        dialogState,
        show,
        hide,
        confirm
    };
}

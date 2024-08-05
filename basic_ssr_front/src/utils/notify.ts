import { Notify } from 'quasar';
import { Dialog } from 'quasar'

export const ConfirmPrompt = (title: any, message: any, actions: any, options: object = {ok:{label: 'ok'}, cancel: {label: 'cancel'}}) => {
  Dialog.create({
    title: title,
    message: message,
    ok: {
      flat: true,
      color: 'primary',
      label: options.ok.label,
    },
    cancel: {
      flat: true,
      color: 'grey',
      label: options.cancel.label
    },
  }).onOk(() => {
    actions()
  })
}

// 错误提示
export const NotifyNegative = (msg: string) => {
  Notify.create({
    type: 'negative',
    position: 'top',
    timeout: 3000,
    message: msg,
  });
};

// 成功提示
export const NotifyPositive = (msg: string) => {
  Notify.create({
    type: 'positive',
    position: 'top',
    timeout: 3000,
    message: msg,
  });
};


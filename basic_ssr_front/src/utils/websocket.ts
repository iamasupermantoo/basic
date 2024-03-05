class WebsocketFunc {
  url: string;                    //  连接路由
  conn: null | any = null;        //  连接对象
  restTime = 5000;      //  重连时间

  onMessageFunc: (msg: any) => void = (() => {/**/})     //  消息回调

  // 构造函数
  constructor(url: string) {
    if (url == '') {
      const protocol = document.location.protocol === 'https:' ? 'wss:' : 'ws:';
      const domain = process.env.baseURL?.substring(process.env.baseURL.indexOf('//'))
      this.url = protocol + domain + '/ws';
    } else {
      this.url = url
    }
  }

  connect() {
    if (this.conn != null) {
      return;
    }

    // 连接websocket
    this.conn = new WebSocket(this.url);

    // 关闭websocket
    this.conn.onclose = () => {
      this.conn = null;

      // 等待时间进行重连
      setTimeout(() => {
        this.connect();
      }, this.restTime);
    };

    // 消息回调
    this.conn.onmessage = (ev: any) => {
      if (this.onMessageFunc != undefined) {
        this.onMessageFunc(JSON.parse(ev.data))
      }
    };
  }

  // 设置消息回调
  setOnMessageFunc(fn: (msg: any) => void) {
    this.onMessageFunc = fn
  }

  // 发送消息
  sendMessageFunc(data: object) {
    if (!this.conn) {
      this.conn.send(JSON.stringify(data));
    }
  }
}

export default WebsocketFunc;

'use strict';

module.exports = (robot) => {
  const pingText = process.env.BENCH_PING_TEXT || 'ping';
  const escapedPingText = pingText.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  const pingPattern = new RegExp(`^\\s*${escapedPingText}\\s*$`, 'i');

  robot.hear(/.*/i, (res) => {
    robot.logger.info(`RUNTIME received text=${JSON.stringify(res.message.text)} room=${JSON.stringify(res.envelope.room)} user=${JSON.stringify(res.envelope.user && res.envelope.user.name)}`);
  });

  const pong = (res) => {
    robot.logger.info(`RUNTIME matched ping text=${JSON.stringify(res.message.text)} room=${JSON.stringify(res.envelope.room)}`);
    res.send('PONG');
    robot.logger.info(`RUNTIME pong sent text=${JSON.stringify(res.message.text)} room=${JSON.stringify(res.envelope.room)}`);
  };

  robot.hear(pingPattern, pong);
  robot.respond(new RegExp(`${escapedPingText}\\s*$`, 'i'), pong);

  const hookDirect = () => {
    if (!robot.direct || robot.__runtimeNotifyHooked) {
      return;
    }
    robot.__runtimeNotifyHooked = true;
    robot.logger.info('RUNTIME direct notify hook installed');
    robot.direct.on('notify_create_message', (msg) => {
      const text = msg && msg.content;
      const room = roomID(msg && msg.talkId);
      robot.logger.info(`RUNTIME direct notify text=${JSON.stringify(text)} room=${JSON.stringify(room)}`);
      if (!room || !pingPattern.test(String(text || ''))) {
        return;
      }
      robot.direct.send({ room }, 'PONG');
      robot.logger.info(`RUNTIME pong sent text=${JSON.stringify(text)} room=${JSON.stringify(room)}`);
    });
  };

  hookDirect();
  robot.logger.info(`READY runtime-node-daab script-loaded direct=${Boolean(robot.direct)}`);
  robot.on('connected', () => {
    robot.logger.info('READY runtime-node-daab connected');
    hookDirect();
  });
};

function roomID(id) {
  if (!id) {
    return '';
  }
  if (typeof id === 'string') {
    return id.startsWith('_') ? id : `_${id}`;
  }
  if (typeof id.high === 'number' && typeof id.low === 'number') {
    return `_${id.high}_${id.low}`;
  }
  if (typeof id.toString === 'function') {
    const text = id.toString();
    return text.startsWith('_') ? text : `_${text}`;
  }
  return '';
}

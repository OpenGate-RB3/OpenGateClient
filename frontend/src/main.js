import Toastify from 'toastify-js'
import 'toastify-js/src/toastify.css'
import {EventsOn} from "../wailsjs/runtime"
import {StartAudioReceive,StartAudioSend,StopAudioReceive,StopAudioSend} from "../wailsjs/go/main/App"


export {StartAudioReceive,StartAudioSend,StopAudioReceive,StopAudioSend}

export function setupEvents(){
    EventsOn("audio_receive_proc",(data)=>{
        if (data === "stopped"){
            Toastify({
                text: "Stop receiving audio",
                duration: 2000,
                gravity: "top",
                position: "right",
                close: true,
                style: {
                    background: "linear-gradient(45deg,rgba(153, 41, 41, 1) 0%, rgba(224, 27, 27, 1) 100%)",
                    color: "white"
                }
            }).showToast()
        }else if(data === "spawned"){
            Toastify({
                text: "Started audio receiving",
                duration: 2000,
                gravity: "top",
                position: "right",
                close: true,
                style: {
                    background: "linear-gradient(180deg, rgba(42, 155, 108, 1) 0%, rgba(14, 171, 105, 1) 50%, rgba(27, 207, 102, 1) 100%)",
                    color: "white"
                }
            }).showToast()
        }
    })
    EventsOn("audio_send_proc",(data)=>{
        if (data === "stopped"){
            Toastify({
                text: "Stop sending audio",
                duration: 3000,
                gravity: "top",
                position: "right",
                close: true,
                style: {
                    background: "linear-gradient(45deg,rgba(153, 41, 41, 1) 0%, rgba(224, 27, 27, 1) 100%)",
                    color: "white"
                }
            }).showToast()
        }else if(data === "spawned"){
            Toastify({
                text: "Started sending audio",
                duration: 3000,
                gravity: "top",
                position: "right",
                close: true,
                style: {
                    background: "linear-gradient(180deg, rgba(42, 155, 108, 1) 0%, rgba(14, 171, 105, 1) 50%, rgba(27, 207, 102, 1) 100%)",
                    color: "white"
                }
            }).showToast()
        }
    })
}
window.setupEvents = setupEvents;
window.StartAudioReceive = StartAudioReceive;
window.StartAudioSend = StartAudioSend;
window.StopAudioReceive = StopAudioReceive;
window.StopAudioSend = StopAudioSend;


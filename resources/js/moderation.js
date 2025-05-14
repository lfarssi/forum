import { showFlashAlert } from "./index.js";

 export function RequestMod(){
    const formReq=document.getElementById("form-requestmod")
    const reason=document.getElementById("reason")
    formReq.addEventListener("submit",async(e)=>{
        e.preventDefault()
         let data = new FormData();
        data.append("reason", reason.value);
        if(reason.value.trim()==""){
                showFlashAlert("reason Is Empty");
        return 
        }else if (reason.value.length<=4){
             showFlashAlert("reason Is too shoort");
        return 
        }

       try {
            const response = await fetch("/reqmod", {
                method: "POST",
                body: data  ,
            });

            if (response.ok) {
                showFlashAlert("your reauest send succ");
                window.location.href = "/"; 
            } else {
                const data = await response.json();
                showFlashAlert(data); 
            }
        } catch (e){
            console.log("error from front"); 
            showFlashAlert(e);
            console.log(e);
            
        }

    })
}
RequestMod()
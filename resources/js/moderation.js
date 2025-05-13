function RequestMod(){
    const formReq=document.querySelector(".form-requestmod")
    const reason=document.querySelector("#reason")
    formReq.addEventListener("submit",async(e)=>{
        e.preventDefault()
         let data = new FormData();
        data.append("reason", content.value);
        if(reason.value.trim()==""){
                showFlashAlert("reason Is Empty");
        return 
        }else if (reason.value.length<=4){
             showFlashAlert("reason Is too shoort");
        return 
        }

       try {
            const response = await fetch("/reqMod", {
                method: "POST",
                body: data,
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
        }

    })
}
<br/>
<div class="forgotmodal-dialog">
    <div class="forgotmodal-container">
        <div id="brand">&nbsp</div>
            &nbsp;
            &nbsp;
            {{if .Verified}}
            <h2>Il tuo account è stato verificato.</h2>
            {{else}}
            <h2>Il tuo account <b>NON</b> è verificato.</h2>
            {{end}}
            </br>
            <div class="check-help">
                <a href="http://{{.domainname}}">Home</a> 
            </div>                          
        </div>
        
    </div>

</div>
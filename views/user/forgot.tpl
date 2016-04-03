<br/>
<div class="forgotmodal-dialog">
    <div class="forgotmodal-container">
        <div id="brand">&nbsp</div>
        &nbsp;
        <h3 class="dark-grey">Recupera Password</h3>
        &nbsp;
        &nbsp;
        {{if .flash.error}}
        <h3>{{.flash.error}}</h3>
        &nbsp;
        {{end}}
        {{if .Errors}}
        {{range $rec := .Errors}}
        <h3>{{$rec}}</h3>
        {{end}}
        &nbsp;
        {{end}}
        {{if .flash.notice}}
        <h3>{{.flash.notice}}</h3>
        &nbsp;
        {{end}}    
        <form method="POST">     
            {{.xsrfdata}}
            <div class="form-group col-lg-12">
                <label>Indirizzo Email</label>
                <input type="email" name="email" class="form-control" placeholder="nome@esempio.com" autofocus>
            </div>   
            <div class="col-sm-12">
                <button type="submit" value="Richiedi Reset" class="btn btn-primary pull-right">Richiedi Reset</button>
            </div>	                       
        </form>
        <div class="check-help">
            <a href="http://{{.domainname}}/">Indietro</a> 
        </div>
    </div>
</div>










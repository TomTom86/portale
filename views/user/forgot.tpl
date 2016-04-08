<br/>
<div class="forgotmodal-dialog">
    <div class="forgotmodal-container">
        <div id="brand">&nbsp</div>
        &nbsp;
        <h3 class="dark-grey">Recupera Password</h3>
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}
            {{if .flash.notice}}
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Operazione completata con successo</strong>
                <hr class="message-inner-separator">
                {{.flash.notice}}.
            </div>
            {{end}}            
            {{if .Errors}}
            {{range $rec := .Errors}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{$rec}}.
            </div>      
            {{end}}
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










<br/>
<div class="forgotmodal-dialog">
    <div class="forgotmodal-container">
        <div id="brand">&nbsp</div>
        &nbsp;
        <h3 class="dark-grey">Reset Password</h3>
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
                <label>Password (min. 6 caratteri)</label>      {{if .Errors.Password}}<b class="error"> ****{{.Errors.Password}}****</b>{{end}}
                <input type="password" name="password" class="form-control"  placeholder="Password">
            </div>
            
            <div class="form-group col-lg-12">
                <label>Ripeti Password</label>      {{if .Errors.Confirm}}<b class="error"><br>****{{.Errors.Confirm}}****</b>{{end}}
                <input type="password" name="password2" class="form-control" placeholder="Password">
            </div>
            <div class="col-sm-12">
                <button type="submit" value="Reset password" class="btn btn-primary pull-right">Reset Password</button>
            </div>	                     
        </form>
        </br>
        <div class="check-help">
            <a href="http://{{.domainname}}">Home</a> 
        </div>                      
    </div>

</div>
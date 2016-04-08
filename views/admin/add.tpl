<div class="col-md-10 content">
        <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Aggiungi Utente</h3>
        </div>
        <div class="panel-body">

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
                {{ .xsrfdata }}   
                    <table>
                    <tr><td><label class="control-label" for="first">Nome: </label></td><td><input name="first" value="{{.User.First}}" size="30" required=""/>{{if .Errors.First}}  {{.Errors.First}}  {{end}}</td></tr>
                    <tr><td><label class="control-label" for="last">Cognome: </label></td><td><input name="last" value="{{.User.Last}}" size="30" required=""/></td></tr>
                    <tr><td><label class="control-label" for="email">Email: </label></td><td><input name="email" value="{{.User.Email}}" size="30" required=""/>{{if .Errors.Email}}  {{.Errors.Email}}  {{end}}</td></tr>
                    <tr><td><label class="control-label" for="password">Password: </label></td><td><input name="password" value="{{.User.Password}}" size="30" required=""/>{{if .Errors.Password}}  {{.Errors.Password}}  {{end}}</td></tr>
                    <tr><td>&nbsp;</td></tr>
                    <tr>
                        <td>&nbsp;</td><td><input type="submit" value="Add" /></td>
                    </tr>
                    </table>
                        <a href="http://{{.domainname}}/admin/{{.parms}}" >Indietro</a>
                </form>
                </div>

        </div>

</div>

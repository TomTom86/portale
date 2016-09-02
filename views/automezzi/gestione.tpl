<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h3>Gestione Utenti</h3>
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
            <div class="form-group  col-lg-2">
                <div class="btn-toolbar pull-left">
                    <button class="btn btn-primary" onclick="location.href='http://{{.domainname}}/admin/add/id!0!id__gte,0'">+ Nuovo Utente</button>
                </div> 
            </div>            

        {{.Ciao}}                 
        </div>
    </div>
</div>
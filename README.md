# Go-Fiber-CRM-Basic

This is a basic CRM system API built with Go-Fiber
<hr>

## Models

### Lead

```golang
type Lead struct {
    gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}
```
<hr>

## Endpoints

|                           <br/> GET <br/><br/>                 |                         POST              |                     PUT                       |                   DELETE                      |
|:--------------------------------------------------------------:|:-----------------------------------------:|:---------------------------------------------:|:---------------------------------------------:|
| <br/>`"api/v1/lead"`<br/>&emsp;* Get all leads<br/><br/>       | `"api/v1/lead"`<br/>&emsp;* Create a lead | `"api/v1/lead/:id"`<br/>&emsp;* Update a lead | `"api/v1/lead/:id"`<br/>&emsp;* Delete a lead |
| <br/>`"api/v1/lead:id"`<br/>&emsp;* Get a lead by id<br/><br/> |                                           |                                               |                                               |

# HongBao
红包的发放和秒杀 后台搭建 通过go实现相关功能 用于个人学习

# 账户表和流水表结构的设计
   ### * 账户表的设计：  
        id  
        account_no  
        account_name  
        account_type
        currency_code  
        user_id  
        user_name  
        balance  
        status  
        created_at  
        updated_at  
        
       
   ### * 流水表的设计  
        id  
        trade_no  
        log_no  
        account_no  
        target_account_no  
        user_id  
        target_user_id  
        amount  
        balance  
        change_type  
        change_flag  
        status  
        decs  
        created_at  
        
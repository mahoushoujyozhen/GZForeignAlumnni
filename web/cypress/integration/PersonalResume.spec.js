// test.spec.js created with Cypress
//
// Start writing your Cypress tests below!
// If you're unfamiliar with how Cypress works,
// check out the link below and learn how to write your first test:
// https://on.cypress.io/writing-first-test

import { each } from "lodash"



describe('login', () => {
    it('visitMenberCenter', () => {

        //访问login界面
        cy.visit('http://localhost:8080/login')
        cy.url().should('include', "/login")

        //输入用户名并检验
        // 用name获取元素
        cy.get("[name='用户名']")
            .type("11")
            .should('have.value', "11")
        // 输入密码并检验
        cy.get("[name='密码']")
            .type("11")
            .should("have.value", "11")
        // 获取按钮，通过text
        cy.get('button:contains("用户登录")').click()
        // 检验是否登录成功，跳转到memberCenter界面
        cy.url().should('include', "memberCenter")
       
        cy.wait(1000)


        // 进入个人简历界面
        cy.get(".funcTitle:contains('个人简历')")
        .click()

        cy.wait(500)

        // 新增工作简历
        cy.url().should("include", "PersonalResueme")
        cy.get('button:contains("新增工作简历")').click()

        // 填写简历信息
        // CSS selector
        cy.get("#van-field-3-input")
        .type('全栈开发')
        .should('have.value','全栈开发')

        cy.get("#van-field-4-input")
        .type('全栈工程师')
        .should('have.value','全栈工程师')

        cy.get("#van-field-5-input")
        .type('软件工程')
        .should('have.value','软件工程')

        cy.get("#van-field-6-input")
        .type('广州大学')
        .should('have.value','广州大学')

        cy.get("#van-field-7-input")
        .type('教育机构')
        .should('have.value','教育机构')

        cy.get("#van-field-8-input")
        .type('18406666666')
        .should('have.value','18406666666')

        cy.get("#van-field-9-input")
        .type('广东省广州市番禺区')
        .should('have.value','广东省广州市番禺区')

        cy.get("#van-field-10-input")
        .click()
        
        cy.get("#app > div.entry > div > div > div > div > div:nth-child(1) > div.van-popup.van-popup--bottom > div > div.van-picker__toolbar > button.van-picker__confirm.van-haptics-feedback")
        .click()

        cy.get("#van-field-11-input")
        .click()
        
        cy.get("#app > div.entry > div > div > div > div > div:nth-child(2) > div.van-popup.van-popup--bottom > div > div.van-picker__toolbar > button.van-picker__confirm.van-haptics-feedback")
        .click()  
        
        cy.wait(1000)

        cy.get("button:contains('提交')")
        .click()

        cy.wait(1000)

        // 新增教育经历

        cy.url().should("include", "PersonalResueme")
        cy.get('button:contains("新增教育简历")')
            .click()
            cy.url().should("include","EdResume")
     
        cy.get("#van-field-12-input")
            .type("本科生")
            .should("have.value","本科生")

            cy.get("#van-field-13-input")
            .type("软件工程")
            .should("have.value","软件工程")

            cy.get("#van-field-14-input")
            .type("广州大学")
            .should("have.value","广州大学")

            cy.get("#van-field-15-input")
            .type("广东省广州市番禺区")
            .should("have.value","广东省广州市番禺区")

            //设置时间
            cy.get("#van-field-16-input")
            .click()
            // cssPath
            cy.get(".van-picker__confirm")
            .click()

            cy.get("#van-field-17-input")
            .click()

            cy.get("div.van-popup:nth-child(10) > div:nth-child(1) > div:nth-child(1) > button:nth-child(2)")
            .click()

            cy.get("div.van-radio:nth-child(2) > div:nth-child(1) > i:nth-child(1)")
            .click()

            // 上传图片
            cy.get("input[type='file']").attachFile('test.png')

            cy.get("button:contains('提交')")
            .click()

            cy.wait(1000)

        // 显示经历详情

        // .entry > div:nth-child(2) > div:nth-child(1)
      cy.contains("全栈开发")
      .click()
      cy.url().should("include","DetailsWork")
      cy.wait(3000)
      cy.go('back')

      cy.contains("本科生")
      .click()
      cy.url().should("include","DetailsEd")
      cy.wait(3000)
      cy.go("back")

      cy.pause()
  
    })

})
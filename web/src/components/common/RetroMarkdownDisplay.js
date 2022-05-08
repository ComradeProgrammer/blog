import React from 'react';
import MarkDownIt from 'markdown-it'
class RetroMarkDownDisplay extends React.Component {

  render() {
    let md = new MarkDownIt({
      html: true,
      linkify: true,
      typographer: true,
    })
    if (!this.props.markdown) {
      return null
    }
    return (
      <div {...this.props} dangerouslySetInnerHTML={{__html: md.render(this.props.markdown)}}>
      </div>
    )
  }
}
export default RetroMarkDownDisplay
/**
 * 模式选择管理 Mixin - 管理深度研究、PPT等模式选择
 */

export default {
  data() {
    return {
      selectedModes: [],
      modeOptions: {
        research: {
          label: '深度研究',
          icon: 'el-icon-aim',
          value: 'research',
          placeholder: '选择一款模型,告诉我想要研究的内容,获取研究报告',
        },
        analysis: {
          label: '数据分析',
          icon: 'el-icon-data-analysis',
          value: 'analysis',
          placeholder: '选择一款模型,上传excel或csv文件,进行数据分析',
        },
        ppt: {
          label: '创建ppt',
          icon: 'el-icon-document',
          value: 'ppt',
          placeholder: '选择一款模型,告诉我想要生成的PPT主题或内容',
        },
        excel: {
          label: '创建excel',
          icon: 'el-icon-s-grid',
          value: 'excel',
          placeholder: '选择一款模型,告诉我想要生成的EXCEL主题或内容',
        },
        web: {
          label: '创建网页',
          icon: 'el-icon-monitor',
          value: 'web',
          placeholder: '选择一款模型,告诉我想要生成的网页主题或内容',
        },
        // video: {
        //   label: '创建视频',
        //   icon: 'el-icon-video-camera',
        //   value: 'video',
        // },
        // skill: {
        //   label: '创建skill',
        //   icon: 'el-icon-cpu',
        //   value: 'skill',
        // },
      },
    };
  },

  methods: {
    /**
     * 添加模式
     */
    addMode(modeValue) {
      // 避免重复添加
      if (this.selectedModes.find(m => m.value === modeValue)) {
        return;
      }
      const mode = this.modeOptions[modeValue];
      if (mode) {
        this.selectedModes.push({ ...mode });
      }
    },

    /**
     * 移除模式
     */
    removeMode(modeValue) {
      this.selectedModes = this.selectedModes.filter(
        m => m.value !== modeValue,
      );
    },

    /**
     * 清空所有模式
     */
    clearModes() {
      this.selectedModes = [];
    },
  },
};
